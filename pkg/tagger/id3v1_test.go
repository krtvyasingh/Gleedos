package tagger

import "testing"

func TestHasID3v1(t *testing.T) {
	data := make([]byte, 128)
	copy(data[0:3], []byte("TAG"))
	if !HasID3v1(data) {
		t.Errorf("expected true for ID3v1 trailer")
	}
}
