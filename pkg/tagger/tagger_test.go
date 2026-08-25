package tagger

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestDetectMediaType(t *testing.T) {
	mp4Header := []byte{0x00, 0x00, 0x00, 0x18, 'f', 't', 'y', 'p', 'm', 'p', '4', '2'}
	mkvHeader := []byte{0x1A, 0x45, 0xDF, 0xA3, 0x93, 0x42, 0x82, 0x88, 'm', 'a', 't', 'r'}
	webmHeader := []byte{0x1A, 0x45, 0xDF, 0xA3, 0x93, 0x42, 0x82, 0x88, 'w', 'e', 'b', 'm'}
	tsHeader := []byte{0x47, 0x40, 0x11, 0x10}

	tests := []struct {
		data     []byte
		expected MediaType
	}{
		{mp4Header, TypeMP4},
		{mkvHeader, TypeMKV},
		{webmHeader, TypeWebM},
		{tsHeader, TypeTS},
		{[]byte("unknown payload"), TypeUnknown},
	}

	for _, tt := range tests {
		got, err := DetectMediaType(bytes.NewReader(tt.data))
		if err != nil {
			t.Errorf("DetectMediaType error: %v", err)
		}
		if got != tt.expected {
			t.Errorf("DetectMediaType = %s, expected %s", got, tt.expected)
		}
	}
}

func TestInjectID3Tag(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gleedos_tagger_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	testMP3 := filepath.Join(tmpDir, "test.mp3")
	dummyAudio := []byte{0xFF, 0xFB, 0x90, 0x64, 0x00, 0x00, 0x00}
	if err := os.WriteFile(testMP3, dummyAudio, 0644); err != nil {
		t.Fatal(err)
	}

	meta := Metadata{
		Title:   "Gleedos Anthem",
		Artist:  "DeepMind",
		Comment: "https://gleedos.local",
	}

	if err := InjectID3Tag(testMP3, meta); err != nil {
		t.Fatalf("InjectID3Tag failed: %v", err)
	}

	mediaType, err := ValidateMediaFile(testMP3)
	if err != nil {
		t.Fatalf("ValidateMediaFile failed: %v", err)
	}
	if mediaType != TypeMP3 {
		t.Fatalf("Expected TypeMP3, got %s", mediaType)
	}
}
