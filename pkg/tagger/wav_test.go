package tagger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsWAV(t *testing.T) {
	tmpDir := t.TempDir()
	wavPath := filepath.Join(tmpDir, "sample.wav")
	hdr := append([]byte("RIFF1234WAVE"), []byte("fmt ")...)
	_ = os.WriteFile(wavPath, hdr, 0644)

	if !IsWAV(wavPath) {
		t.Errorf("expected true for WAV RIFF header")
	}
}
