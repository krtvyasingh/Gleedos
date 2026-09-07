package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSafeWriteAtomic(t *testing.T) {
	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "file.txt")
	err := SafeWriteAtomic(target, []byte("atomic_data"))
	if err != nil {
		t.Fatalf("SafeWriteAtomic failed: %v", err)
	}
	read, _ := os.ReadFile(target)
	if string(read) != "atomic_data" {
		t.Errorf("data mismatch")
	}
}
