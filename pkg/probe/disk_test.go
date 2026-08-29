package probe

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	nested := filepath.Join(tmpDir, "a", "b", "c")
	if err := EnsureDirectory(nested); err != nil {
		t.Fatalf("EnsureDirectory failed: %v", err)
	}
	if _, err := os.Stat(nested); err != nil {
		t.Errorf("directory was not created: %v", err)
	}
}
