package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/krtvysingh/gleedos/pkg/chunker"
	"github.com/krtvysingh/gleedos/pkg/hls"
	"github.com/krtvysingh/gleedos/pkg/limiter"
)

func runDownload(strategy downloadStrategy) error {
	cmd := exec.Command("yt-dlp", strategy.args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// sanitizeFilename strips dangerous characters and path separators from filenames.
func sanitizeFilename(name string) string {
	name = filepath.Base(name)
	name = strings.TrimSpace(name)
	if name == "" || name == "." || name == "/" || name == "\\" {
		return "download"
	}

	// Remove control characters and characters unsafe on Unix/Windows
	reg := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)
	cleaned := reg.ReplaceAllString(name, "_")
	cleaned = strings.Trim(cleaned, ". ")
	if cleaned == "" {
		return "download"
	}
	return cleaned
}

// isDirectNativeURL checks if the URL is a direct media file or HLS stream suitable for zero-dependency download.
func isDirectNativeURL(u string) (bool, string) {
	parsed, err := url.Parse(u)
	if err != nil {
		return false, ""
	}

	path := strings.ToLower(parsed.Path)
	if strings.HasSuffix(path, ".m3u8") {
		return true, "hls"
	}

	directExts := []string{".mp4", ".mkv", ".webm", ".mp3", ".flac", ".aac", ".wav", ".ts"}
	for _, ext := range directExts {
		if strings.HasSuffix(path, ext) {
			return true, ext[1:]
		}
	}

	return false, ""
}

// runNativeDownload downloads media using pure Go without external dependencies.
func runNativeDownload(
	ctx context.Context,
	mediaURL string,
	outputDir string,
	kind string,
	concurrency int,
	rateLim *limiter.RateLimiter,
) (string, error) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("create output directory: %w", err)
	}

	parsed, _ := url.Parse(mediaURL)
	rawName := filepath.Base(parsed.Path)
	filename := sanitizeFilename(rawName)

	if kind == "hls" && !strings.HasSuffix(strings.ToLower(filename), ".ts") && !strings.HasSuffix(strings.ToLower(filename), ".mp4") {
		filename = strings.TrimSuffix(filename, filepath.Ext(filename)) + ".ts"
	}

	targetPath := filepath.Join(outputDir, filename)

	if kind == "hls" {
		hlsDownloader := hls.New(hls.Config{
			URL:         mediaURL,
			TargetPath:  targetPath,
			Concurrency: concurrency,
			RateLimiter: rateLim,
		})
		if err := hlsDownloader.Download(ctx); err != nil {
			return "", err
		}
		return targetPath, nil
	}

	chunkDownloader := chunker.New(chunker.Config{
		URL:         mediaURL,
		TargetPath:  targetPath,
		Concurrency: concurrency,
		RateLimiter: rateLim,
	})

	if err := chunkDownloader.Download(ctx); err != nil {
		return "", err
	}

	return targetPath, nil
}
