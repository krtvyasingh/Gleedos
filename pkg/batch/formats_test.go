package batch

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePlaylistFile(t *testing.T) {
	tmpDir := t.TempDir()
	jsonPath := filepath.Join(tmpDir, "list.json")
	_ = os.WriteFile(jsonPath, []byte(`{"name": "test", "urls": ["https://example.com/1", "https://example.com/2"]}`), 0644)

	urls, err := ParsePlaylistFile(jsonPath)
	if err != nil || len(urls) != 2 {
		t.Fatalf("ParsePlaylistFile failed: %v, got %d urls", err, len(urls))
	}
}
