package vvc

import "testing"

func TestCodecIdentification(t *testing.T) {
	if !IsAV1([]byte("...av01...")) {
		t.Errorf("expected AV1 detection")
	}
	if !IsVVC([]byte("...vvc1...")) {
		t.Errorf("expected VVC detection")
	}
}
