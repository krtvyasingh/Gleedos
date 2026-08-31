package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCleanStaleTempFiles(t *testing.T) {
	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "old.tmp")
	_ = os.WriteFile(f, []byte("1"), 0644)
	past := time.Now().Add(-2 * time.Hour)
	_ = os.Chtimes(f, past, past)

	cleaned := CleanStaleTempFiles(tmpDir, 1*time.Hour)
	if cleaned != 1 {
		t.Errorf("expected 1 file cleaned, got %d", cleaned)
	}
}
