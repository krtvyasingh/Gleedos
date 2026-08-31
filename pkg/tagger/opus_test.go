package tagger

import "testing"

func TestIsOpusHead(t *testing.T) {
	if !IsOpusHead([]byte("OpusHead1234")) {
		t.Errorf("expected true for OpusHead")
	}
}
