package tagger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsOgg(t *testing.T) {
	tmpDir := t.TempDir()
	oggPath := filepath.Join(tmpDir, "audio.ogg")
	_ = os.WriteFile(oggPath, []byte("OggS01234567"), 0644)

	if !IsOgg(oggPath) {
		t.Errorf("expected true for Ogg header")
	}
}
