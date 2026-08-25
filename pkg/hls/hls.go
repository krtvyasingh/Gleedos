package hls

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/krtvysingh/gleedos/pkg/limiter"
	"github.com/krtvysingh/gleedos/pkg/ui"
)

// Segment represents a single media segment in an HLS stream.
type Segment struct {
	Index    int
	Duration float64
	URL      string
}

// Playlist represents parsed HLS playlist data.
type Playlist struct {
	IsMaster bool
	Variants []string
	Segments []Segment
}

// ParsePlaylist parses M3U8 playlist content.
func ParsePlaylist(baseURL *url.URL, r io.Reader) (*Playlist, error) {
	scanner := bufio.NewScanner(r)
	p := &Playlist{}

	var currentDuration float64
	segIdx := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#EXT-X-STREAM-INF") {
			p.IsMaster = true
			continue
		}

		if strings.HasPrefix(line, "#EXTINF:") {
			parts := strings.Split(strings.TrimPrefix(line, "#EXTINF:"), ",")
			if len(parts) > 0 {
				dur, _ := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
				currentDuration = dur
			}
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		// Line is a URI
		parsedURI, err := url.Parse(line)
		if err != nil {
			continue
		}
		resolved := baseURL.ResolveReference(parsedURI).String()

		if p.IsMaster {
			p.Variants = append(p.Variants, resolved)
		} else {
			p.Segments = append(p.Segments, Segment{
				Index:    segIdx,
				Duration: currentDuration,
				URL:      resolved,
			})
			segIdx++
			currentDuration = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return p, nil
}

// Config specifies HLS downloader parameters.
type Config struct {
	URL         string
	TargetPath  string
	Concurrency int
	RateLimiter *limiter.RateLimiter
	Headers     map[string]string
	Client      *http.Client
	Quiet       bool
}

// Downloader manages HLS streaming media extraction.
type Downloader struct {
	cfg    Config
	client *http.Client
}

// New creates a new HLS Downloader.
func New(cfg Config) *Downloader {
	if cfg.Concurrency <= 0 {
		cfg.Concurrency = 5
	}
	if cfg.Client == nil {
		cfg.Client = &http.Client{Timeout: 30 * time.Second}
	}
	return &Downloader{cfg: cfg, client: cfg.Client}
}

// Download fetches and merges all segments from the M3U8 stream.
func (d *Downloader) Download(ctx context.Context) error {
	playlistURL := d.cfg.URL

	// 1. Fetch Master or Media Playlist
	playlist, resolvedURL, err := d.fetchPlaylist(ctx, playlistURL)
	if err != nil {
		return fmt.Errorf("fetch playlist: %w", err)
	}

	if playlist.IsMaster {
		if len(playlist.Variants) == 0 {
			return errors.New("master playlist contains no variants")
		}
		// Pick highest bandwidth / last variant
		variantURL := playlist.Variants[len(playlist.Variants)-1]
		playlist, _, err = d.fetchPlaylist(ctx, variantURL)
		if err != nil {
			return fmt.Errorf("fetch variant playlist: %w", err)
		}
	}

	if len(playlist.Segments) == 0 {
		return errors.New("no media segments found in playlist")
	}

	_ = resolvedURL

	// 2. Prepare staging directory for temporary segments
	stagingDir := d.cfg.TargetPath + ".hls_parts"
	if err := os.MkdirAll(stagingDir, 0755); err != nil {
		return fmt.Errorf("mkdir staging: %w", err)
	}
	defer os.RemoveAll(stagingDir)

	totalSegments := len(playlist.Segments)
	tracker := ui.NewProgressTracker(int64(totalSegments), d.cfg.Concurrency)
	tracker.SetStatus("Downloading HLS")
	tracker.SetQuiet(d.cfg.Quiet)

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			tracker.PrintProgress()
		}
	}()

	// 3. Download segments in parallel worker pool
	segChan := make(chan Segment, totalSegments)
	for _, seg := range playlist.Segments {
		segChan <- seg
	}
	close(segChan)

	var wg sync.WaitGroup
	errChan := make(chan error, d.cfg.Concurrency)

	workerCount := d.cfg.Concurrency
	if workerCount > totalSegments {
		workerCount = totalSegments
	}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for seg := range segChan {
				select {
				case <-ctx.Done():
					errChan <- ctx.Err()
					return
				default:
				}

				partPath := filepath.Join(stagingDir, fmt.Sprintf("seg_%06d.ts", seg.Index))
				if err := d.downloadSegment(ctx, seg.URL, partPath); err != nil {
					errChan <- fmt.Errorf("seg %d: %w", seg.Index, err)
					return
				}
				tracker.AddBytes(1)
			}
		}()
	}

	wg.Wait()
	close(errChan)

	if err, ok := <-errChan; ok && err != nil {
		return err
	}

	tracker.PrintProgress()
	tracker.Finish()

	// 4. Merge all segment parts sequentially into TargetPath
	outFile, err := os.Create(d.cfg.TargetPath)
	if err != nil {
		return fmt.Errorf("create target: %w", err)
	}
	defer outFile.Close()

	for i := 0; i < totalSegments; i++ {
		partPath := filepath.Join(stagingDir, fmt.Sprintf("seg_%06d.ts", i))
		partFile, err := os.Open(partPath)
		if err != nil {
			return fmt.Errorf("open segment %d: %w", i, err)
		}
		_, copyErr := io.Copy(outFile, partFile)
		partFile.Close()
		if copyErr != nil {
			return fmt.Errorf("merge segment %d: %w", i, copyErr)
		}
	}

	return nil
}

func (d *Downloader) fetchPlaylist(ctx context.Context, u string) (*Playlist, *url.URL, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range d.cfg.Headers {
		req.Header.Set(k, v)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("status %s", resp.Status)
	}

	base, err := url.Parse(u)
	if err != nil {
		return nil, nil, err
	}

	p, err := ParsePlaylist(base, resp.Body)
	return p, base, err
}

func (d *Downloader) downloadSegment(ctx context.Context, segURL, outPath string) error {
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, "GET", segURL, nil)
		if err != nil {
			return err
		}
		for k, v := range d.cfg.Headers {
			req.Header.Set(k, v)
		}

		resp, err := d.client.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			lastErr = fmt.Errorf("status %s", resp.Status)
			time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
			continue
		}

		f, err := os.Create(outPath)
		if err != nil {
			resp.Body.Close()
			return err
		}

		r := limiter.NewReader(ctx, resp.Body, d.cfg.RateLimiter)
		_, copyErr := io.Copy(f, r)
		f.Close()
		resp.Body.Close()

		if copyErr == nil {
			return nil
		}
		lastErr = copyErr
	}

	return lastErr
}
