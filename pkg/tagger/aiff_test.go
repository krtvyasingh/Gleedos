package tagger

import "testing"

func TestIsAIFF(t *testing.T) {
	hdr := []byte("FORM1234AIFF")
	if !IsAIFF(hdr) {
		t.Errorf("expected true for AIFF header")
	}
}
