package batch

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestParseURLFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gleedos_batch_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	batchFile := filepath.Join(tmpDir, "urls.txt")
	content := `# Video list
https://example.com/video1

https://example.com/video2
# Comment line
https://example.com/video3
`
	if err := os.WriteFile(batchFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	urls, err := ParseURLFile(batchFile)
	if err != nil {
		t.Fatalf("ParseURLFile failed: %v", err)
	}

	if len(urls) != 3 {
		t.Fatalf("expected 3 urls, got %d", len(urls))
	}
}

func TestProcessBatch(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gleedos_batch_proc_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	historyFile := filepath.Join(tmpDir, "history.json")
	store, err := NewHistoryStore(historyFile)
	if err != nil {
		t.Fatal(err)
	}

	urls := []string{
		"https://example.com/1",
		"https://example.com/2",
		"https://example.com/3",
	}

	success, fail, err := ProcessBatch(
		context.Background(),
		urls,
		2,
		store,
		func(ctx context.Context, u string) (string, error) {
			return "/downloads/output.mp4", nil
		},
	)

	if err != nil {
		t.Fatalf("ProcessBatch failed: %v", err)
	}
	if success != 3 || fail != 0 {
		t.Fatalf("expected 3 successes 0 fails, got %d s / %d f", success, fail)
	}

	// Verify store has URLs
	if !store.HasURL("https://example.com/1") {
		t.Errorf("expected URL 1 in history store")
	}
}
