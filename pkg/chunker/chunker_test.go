package chunker

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func createMockRangeServer(data []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept-Ranges", "bytes")
		total := len(data)

		if r.Method == "HEAD" {
			w.Header().Set("Content-Length", strconv.Itoa(total))
			w.WriteHeader(http.StatusOK)
			return
		}

		rangeHdr := r.Header.Get("Range")
		if rangeHdr == "" {
			w.Header().Set("Content-Length", strconv.Itoa(total))
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}

		// Handle bytes=start-end
		if strings.HasPrefix(rangeHdr, "bytes=") {
			parts := strings.Split(strings.TrimPrefix(rangeHdr, "bytes="), "-")
			start, _ := strconv.Atoi(parts[0])
			end := total - 1
			if len(parts) > 1 && parts[1] != "" {
				end, _ = strconv.Atoi(parts[1])
			}
			if end >= total {
				end = total - 1
			}

			w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, total))
			w.Header().Set("Content-Length", strconv.Itoa(end-start+1))
			w.WriteHeader(http.StatusPartialContent)
			w.Write(data[start : end+1])
			return
		}

		http.Error(w, "invalid range", http.StatusBadRequest)
	}))
}

func TestChunkerDownload(t *testing.T) {
	// Create 2MB test payload
	payload := make([]byte, 2*1024*1024)
	for i := range payload {
		payload[i] = byte(i % 256)
	}

	server := createMockRangeServer(payload)
	defer server.Close()

	tmpDir, err := os.MkdirTemp("", "gleedos_test_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	targetFile := filepath.Join(tmpDir, "test_output.bin")

	downloader := New(Config{
		URL:         server.URL,
		TargetPath:  targetFile,
		Concurrency: 4,
		Quiet:       true,
	})

	err = downloader.Download(context.Background())
	if err != nil {
		t.Fatalf("Download failed: %v", err)
	}

	// Verify downloaded content
	result, err := os.ReadFile(targetFile)
	if err != nil {
		t.Fatalf("Read target file failed: %v", err)
	}

	if !bytes.Equal(result, payload) {
		t.Fatalf("Downloaded payload mismatch! expected %d bytes, got %d", len(payload), len(result))
	}

	// Verify .meta file was cleaned up on successful completion
	metaPath := targetFile + ".gleedos.meta"
	if _, err := os.Stat(metaPath); !os.IsNotExist(err) {
		t.Errorf("expected .meta file to be cleaned up, but it still exists")
	}
}

func TestChunkerResume(t *testing.T) {
	payload := make([]byte, 1024*1024)
	for i := range payload {
		payload[i] = byte((i * 7) % 256)
	}

	server := createMockRangeServer(payload)
	defer server.Close()

	tmpDir, err := os.MkdirTemp("", "gleedos_resume_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	targetFile := filepath.Join(tmpDir, "test_resume.bin")

	// Simulate first partial attempt: cancel context halfway
	ctx, cancel := context.WithCancel(context.Background())
	downloader1 := New(Config{
		URL:         server.URL,
		TargetPath:  targetFile,
		Concurrency: 4,
		Quiet:       true,
	})

	// Cancel shortly after starting
	go func() {
		cancel()
	}()

	_ = downloader1.Download(ctx)

	// Now resume with a fresh context
	downloader2 := New(Config{
		URL:         server.URL,
		TargetPath:  targetFile,
		Concurrency: 4,
		Quiet:       true,
	})

	err = downloader2.Download(context.Background())
	if err != nil {
		t.Fatalf("Resumed download failed: %v", err)
	}

	result, err := os.ReadFile(targetFile)
	if err != nil {
		t.Fatalf("Read resumed file failed: %v", err)
	}

	if !bytes.Equal(result, payload) {
		t.Fatalf("Resumed payload mismatch! expected %d bytes, got %d", len(payload), len(result))
	}
}
