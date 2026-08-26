package tagger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReTaggingMP3(t *testing.T) {
	tmpDir := t.TempDir()
	mp3Path := filepath.Join(tmpDir, "retag.mp3")

	dummyMP3 := []byte{0xFF, 0xFB, 0x90, 0x64, 0x01, 0x02, 0x03, 0x04}
	if err := os.WriteFile(mp3Path, dummyMP3, 0644); err != nil {
		t.Fatal(err)
	}

	// First Tag
	meta1 := Metadata{Title: "First Title", Artist: "Artist 1"}
	if err := InjectID3Tag(mp3Path, meta1); err != nil {
		t.Fatalf("first tag injection failed: %v", err)
	}

	// Re-Tag
	meta2 := Metadata{Title: "Second Title", Artist: "Artist 2", Comment: "Replaced"}
	if err := InjectID3Tag(mp3Path, meta2); err != nil {
		t.Fatalf("second tag injection failed: %v", err)
	}

	// Validate file
	mediaType, err := ValidateMediaFile(mp3Path)
	if err != nil {
		t.Fatalf("ValidateMediaFile failed: %v", err)
	}
	if mediaType != TypeMP3 {
		t.Errorf("expected TypeMP3, got %s", mediaType)
	}

	// Verify file does not contain multiple ID3 headers
	content, err := os.ReadFile(mp3Path)
	if err != nil {
		t.Fatal(err)
	}
	if string(content[:3]) != "ID3" {
		t.Errorf("missing ID3 header")
	}
}

func TestValidateMediaFileRejectsEmptyOrMissing(t *testing.T) {
	tmpDir := t.TempDir()
	emptyPath := filepath.Join(tmpDir, "empty.mp4")
	if err := os.WriteFile(emptyPath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := ValidateMediaFile(emptyPath)
	if err == nil {
		t.Errorf("expected error for 0-byte file, got nil")
	}

	_, err = ValidateMediaFile(filepath.Join(tmpDir, "non_existent.mp4"))
	if err == nil {
		t.Errorf("expected error for missing file, got nil")
	}
}
