package tagger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsFLAC(t *testing.T) {
	tmpDir := t.TempDir()
	flacPath := filepath.Join(tmpDir, "sample.flac")
	_ = os.WriteFile(flacPath, []byte("fLaC1234"), 0644)

	if !IsFLAC(flacPath) {
		t.Errorf("expected true for FLAC header")
	}
}
