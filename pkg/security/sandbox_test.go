package security

import (
	"os"
	"runtime"
	"testing"
)

func TestCreateSandboxDir(t *testing.T) {
	tmpDir := t.TempDir()
	sb, err := CreateSandboxDir(tmpDir)
	if err != nil {
		t.Fatalf("CreateSandboxDir failed: %v", err)
	}
	stat, err := os.Stat(sb)
	if err != nil {
		t.Fatalf("os.Stat failed: %v", err)
	}
	if runtime.GOOS != "windows" && stat.Mode().Perm() != 0700 {
		t.Errorf("expected 0700 permissions, got %v", stat.Mode().Perm())
	}
}
