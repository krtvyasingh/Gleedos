package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAtomic(t *testing.T) {
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "atomic.dat")
	if err := WriteAtomic(file, []byte("data")); err != nil {
		t.Fatalf("WriteAtomic failed: %v", err)
	}
	if _, err := os.Stat(file); err != nil {
		t.Errorf("file not found")
	}
}
