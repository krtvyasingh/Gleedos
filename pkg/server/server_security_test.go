package server

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestServerSecurityRejections(t *testing.T) {
	srv := New(8998, func(ctx context.Context, req DownloadRequest) (string, error) {
		return "/tmp/out.mp4", nil
	}, nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_ = srv.Start(ctx)
	}()

	time.Sleep(100 * time.Millisecond)

	// 1. Method Not Allowed on /api/download with GET
	resp, err := http.Get("http://127.0.0.1:8998/api/download")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected 405 Method Not Allowed, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// 2. Reject non-HTTP URL (SSRF defense)
	badPayload := `{"url": "file:///etc/passwd"}`
	resp, err = http.Post("http://127.0.0.1:8998/api/download", "application/json", strings.NewReader(badPayload))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request for file:// URL, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// 3. Reject oversized request payload (> 1MB)
	hugePayload := bytes.Repeat([]byte("a"), 2*1024*1024)
	resp, err = http.Post("http://127.0.0.1:8998/api/download", "application/json", bytes.NewReader(hugePayload))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest && resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("expected 400/413 for oversized body, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// 4. Missing ID parameter on status
	resp, err = http.Get("http://127.0.0.1:8998/api/status")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 for missing id, got %d", resp.StatusCode)
	}
	resp.Body.Close()
}
