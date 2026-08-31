package security

import (
	"os"
	"testing"
)

func TestCreateSandboxDir(t *testing.T) {
	tmpDir := t.TempDir()
	sb, err := CreateSandboxDir(tmpDir)
	if err != nil {
		t.Fatalf("CreateSandboxDir failed: %v", err)
	}
	stat, _ := os.Stat(sb)
	if stat.Mode().Perm() != 0700 {
		t.Errorf("expected 0700 permissions, got %v", stat.Mode().Perm())
	}
}
