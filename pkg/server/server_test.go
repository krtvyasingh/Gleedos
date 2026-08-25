package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/krtvysingh/gleedos/pkg/batch"
)

func TestServerEndpoints(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gleedos_srv_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	store, err := batch.NewHistoryStore(filepath.Join(tmpDir, "hist.json"))
	if err != nil {
		t.Fatal(err)
	}

	handler := func(ctx context.Context, req DownloadRequest) (string, error) {
		time.Sleep(50 * time.Millisecond)
		return "/tmp/downloaded_file.mp4", nil
	}

	srv := New(8999, handler, store)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_ = srv.Start(ctx)
	}()

	time.Sleep(100 * time.Millisecond) // Let server start

	// 1. Health check
	resp, err := http.Get("http://127.0.0.1:8999/health")
	if err != nil {
		t.Fatalf("Health check failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK from health, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// 2. Submit download
	reqBody := `{"url": "https://example.com/video"}`
	resp, err = http.Post("http://127.0.0.1:8999/api/download", "application/json", bytes.NewReader([]byte(reqBody)))
	if err != nil {
		t.Fatalf("Post download failed: %v", err)
	}
	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("Expected 202 Accepted, got %d", resp.StatusCode)
	}

	var job JobResponse
	if err := json.NewDecoder(resp.Body).Decode(&job); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	resp.Body.Close()

	// 3. Poll status
	time.Sleep(150 * time.Millisecond)
	resp, err = http.Get("http://127.0.0.1:8999/api/status?id=" + job.ID)
	if err != nil {
		t.Fatalf("Get status failed: %v", err)
	}
	var pollJob JobResponse
	json.NewDecoder(resp.Body).Decode(&pollJob)
	resp.Body.Close()

	if pollJob.Status != StatusCompleted {
		t.Fatalf("Expected status completed, got %s", pollJob.Status)
	}
	if pollJob.OutputPath != "/tmp/downloaded_file.mp4" {
		t.Errorf("Unexpected output path: %s", pollJob.OutputPath)
	}
}
