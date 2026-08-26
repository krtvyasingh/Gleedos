package chunker

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/krtvysingh/gleedos/pkg/limiter"
	"github.com/krtvysingh/gleedos/pkg/ui"
)

// Chunk represents a discrete byte range in a download.
type Chunk struct {
	Index      int   `json:"index"`
	Start      int64 `json:"start"`
	End        int64 `json:"end"`
	Downloaded int64 `json:"downloaded"`
	Completed  bool  `json:"completed"`
}

// DownloadMeta contains checkpoint data for resuming downloads.
type DownloadMeta struct {
	URL           string  `json:"url"`
	ContentLength int64   `json:"content_length"`
	TotalChunks   int     `json:"total_chunks"`
	Chunks        []Chunk `json:"chunks"`
	ETag          string  `json:"etag"`
	LastModified  string  `json:"last_modified"`
}

// Config specifies chunker download options.
type Config struct {
	URL         string
	TargetPath  string
	Concurrency int
	RateLimiter *limiter.RateLimiter
	Headers     map[string]string
	Client      *http.Client
	Quiet       bool
}

// Downloader manages parallel chunked downloading.
type Downloader struct {
	cfg      Config
	client   *http.Client
	metaPath string
	meta     *DownloadMeta
	mu       sync.Mutex
	tracker  *ui.ProgressTracker
}

// defaultTransport provides secure and robust connection timeouts.
func defaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}
}

// New creates a new Downloader.
func New(cfg Config) *Downloader {
	if cfg.Concurrency <= 0 {
		cfg.Concurrency = 4
	}
	if cfg.Concurrency > 32 {
		cfg.Concurrency = 32 // Prevent excessive connection abuse
	}
	if cfg.Client == nil {
		cfg.Client = &http.Client{
			Transport: defaultTransport(),
		}
	}
	metaPath := cfg.TargetPath + ".gleedos.meta"
	return &Downloader{
		cfg:      cfg,
		client:   cfg.Client,
		metaPath: metaPath,
	}
}

// validateURL ensures the URL is syntactically valid and uses HTTP/HTTPS.
func validateURL(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid URL syntax: %w", err)
	}
	scheme := strings.ToLower(u.Scheme)
	if scheme != "http" && scheme != "https" {
		return fmt.Errorf("unsupported protocol scheme: %s (only http and https allowed)", u.Scheme)
	}
	if u.Host == "" {
		return errors.New("URL missing host")
	}
	return nil
}

// Download starts or resumes the parallel chunked download.
func (d *Downloader) Download(ctx context.Context) error {
	if err := validateURL(d.cfg.URL); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "HEAD", d.cfg.URL, nil)
	if err != nil {
		return fmt.Errorf("create HEAD request: %w", err)
	}
	for k, v := range d.cfg.Headers {
		req.Header.Set(k, v)
	}

	resp, err := d.client.Do(req)
	var probeResp *http.Response

	if err != nil || resp.StatusCode >= 400 {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
		// Try GET with Range 0-0 fallback
		probeReq, pErr := http.NewRequestWithContext(ctx, "GET", d.cfg.URL, nil)
		if pErr != nil {
			return fmt.Errorf("create probe GET request: %w", pErr)
		}
		for k, v := range d.cfg.Headers {
			probeReq.Header.Set(k, v)
		}
		probeReq.Header.Set("Range", "bytes=0-0")
		probeResp, err = d.client.Do(probeReq)
		if err != nil {
			return fmt.Errorf("probe request failed: %w", err)
		}
		resp = probeResp
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("server returned unexpected status: %s", resp.Status)
	}

	acceptRanges := resp.Header.Get("Accept-Ranges") == "bytes" || resp.StatusCode == http.StatusPartialContent
	contentLength := resp.ContentLength
	etag := resp.Header.Get("ETag")
	lastMod := resp.Header.Get("Last-Modified")

	if cr := resp.Header.Get("Content-Range"); cr != "" {
		var start, end, total int64
		if _, scanErr := fmt.Sscanf(cr, "bytes %d-%d/%d", &start, &end, &total); scanErr == nil && total > 0 {
			contentLength = total
			acceptRanges = true
		}
	}

	// Ensure destination directory exists
	targetDir := filepath.Dir(d.cfg.TargetPath)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("mkdir target dir: %w", err)
	}

	// Fallback to single-stream download if ranges not supported or length unknown
	if !acceptRanges || contentLength <= 0 || d.cfg.Concurrency <= 1 {
		return d.downloadSingleStream(ctx)
	}

	return d.downloadParallel(ctx, contentLength, etag, lastMod)
}

func (d *Downloader) downloadSingleStream(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", d.cfg.URL, nil)
	if err != nil {
		return err
	}
	for k, v := range d.cfg.Headers {
		req.Header.Set(k, v)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("single stream status: %s", resp.Status)
	}

	tmpFile := d.cfg.TargetPath + ".part"
	outFile, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	bodyReader := limiter.NewReader(ctx, resp.Body, d.cfg.RateLimiter)
	tracker := ui.NewProgressTracker(resp.ContentLength, 1)
	tracker.SetQuiet(d.cfg.Quiet)

	buf := make([]byte, 32*1024)
	var written int64
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			tracker.PrintProgress()
		}
	}()

	for {
		n, rErr := bodyReader.Read(buf)
		if n > 0 {
			wN, wErr := outFile.Write(buf[:n])
			if wErr != nil {
				return wErr
			}
			written += int64(wN)
			tracker.AddBytes(int64(wN))
		}
		if rErr != nil {
			if errors.Is(rErr, io.EOF) {
				break
			}
			return rErr
		}
	}

	tracker.PrintProgress()
	tracker.Finish()
	_ = outFile.Close()

	return os.Rename(tmpFile, d.cfg.TargetPath)
}

func (d *Downloader) loadOrCreateMeta(contentLength int64, etag, lastMod string) error {
	if data, err := os.ReadFile(d.metaPath); err == nil {
		var meta DownloadMeta
		if json.Unmarshal(data, &meta) == nil {
			if meta.URL == d.cfg.URL && meta.ContentLength == contentLength {
				d.meta = &meta
				return nil
			}
		}
	}

	// Create new chunk plan
	chunkSize := contentLength / int64(d.cfg.Concurrency)
	if chunkSize < 512*1024 {
		chunkSize = 512 * 1024
	}

	chunks := make([]Chunk, 0)
	var start int64
	idx := 0
	for start < contentLength {
		end := start + chunkSize - 1
		if end >= contentLength {
			end = contentLength - 1
		}
		chunks = append(chunks, Chunk{
			Index:      idx,
			Start:      start,
			End:        end,
			Downloaded: 0,
			Completed:  false,
		})
		start = end + 1
		idx++
	}

	d.meta = &DownloadMeta{
		URL:           d.cfg.URL,
		ContentLength: contentLength,
		TotalChunks:   len(chunks),
		Chunks:        chunks,
		ETag:          etag,
		LastModified:  lastMod,
	}

	return d.saveMeta()
}

func (d *Downloader) saveMeta() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	data, err := json.Marshal(d.meta)
	if err != nil {
		return err
	}
	tmpMeta := d.metaPath + ".tmp"
	if err := os.WriteFile(tmpMeta, data, 0600); err != nil {
		return err
	}
	return os.Rename(tmpMeta, d.metaPath)
}

func (d *Downloader) downloadParallel(ctx context.Context, contentLength int64, etag, lastMod string) error {
	if err := d.loadOrCreateMeta(contentLength, etag, lastMod); err != nil {
		return fmt.Errorf("init meta: %w", err)
	}

	file, err := os.OpenFile(d.cfg.TargetPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("open target file: %w", err)
	}
	defer file.Close()

	if stat, err := file.Stat(); err == nil && stat.Size() < contentLength {
		if tErr := file.Truncate(contentLength); tErr != nil {
			return fmt.Errorf("pre-allocate file space: %w", tErr)
		}
	}

	var totalInitialDownloaded int64
	for _, ch := range d.meta.Chunks {
		totalInitialDownloaded += ch.Downloaded
	}

	tracker := ui.NewProgressTracker(contentLength, d.cfg.Concurrency)
	tracker.SetQuiet(d.cfg.Quiet)
	tracker.AddBytes(totalInitialDownloaded)
	d.tracker = tracker

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			tracker.PrintProgress()
		}
	}()

	chunkChan := make(chan Chunk, len(d.meta.Chunks))
	for _, ch := range d.meta.Chunks {
		if !ch.Completed {
			chunkChan <- ch
		}
	}
	close(chunkChan)

	var wg sync.WaitGroup
	errChan := make(chan error, d.cfg.Concurrency)

	workerCount := d.cfg.Concurrency
	if workerCount > len(d.meta.Chunks) {
		workerCount = len(d.meta.Chunks)
	}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ch := range chunkChan {
				select {
				case <-ctx.Done():
					errChan <- ctx.Err()
					return
				default:
				}

				if err := d.downloadChunk(ctx, file, ch); err != nil {
					errChan <- err
					return
				}
			}
		}()
	}

	wg.Wait()
	close(errChan)

	if err, ok := <-errChan; ok && err != nil {
		_ = d.saveMeta()
		return err
	}

	tracker.PrintProgress()
	tracker.Finish()

	allDone := true
	for _, ch := range d.meta.Chunks {
		if !ch.Completed {
			allDone = false
			break
		}
	}

	if allDone {
		_ = os.Remove(d.metaPath)
	}

	return nil
}

func (d *Downloader) downloadChunk(ctx context.Context, file *os.File, ch Chunk) error {
	fetchStart := ch.Start + ch.Downloaded
	if fetchStart > ch.End {
		d.mu.Lock()
		d.meta.Chunks[ch.Index].Completed = true
		d.mu.Unlock()
		return nil
	}

	req, err := http.NewRequestWithContext(ctx, "GET", d.cfg.URL, nil)
	if err != nil {
		return err
	}
	for k, v := range d.cfg.Headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", fetchStart, ch.End))

	resp, err := d.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("chunk %d HTTP status %s", ch.Index, resp.Status)
	}

	bodyReader := limiter.NewReader(ctx, resp.Body, d.cfg.RateLimiter)
	buf := make([]byte, 32*1024)
	currentOffset := fetchStart

	for {
		n, rErr := bodyReader.Read(buf)
		if n > 0 {
			if _, wErr := file.WriteAt(buf[:n], currentOffset); wErr != nil {
				return fmt.Errorf("write at offset %d: %w", currentOffset, wErr)
			}
			currentOffset += int64(n)

			d.mu.Lock()
			d.meta.Chunks[ch.Index].Downloaded += int64(n)
			d.mu.Unlock()

			d.tracker.AddBytes(int64(n))
		}

		if rErr != nil {
			if errors.Is(rErr, io.EOF) {
				break
			}
			return rErr
		}
	}

	d.mu.Lock()
	d.meta.Chunks[ch.Index].Completed = true
	d.mu.Unlock()

	_ = d.saveMeta()
	return nil
}
