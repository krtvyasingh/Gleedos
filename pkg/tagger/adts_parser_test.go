package tagger

import "testing"

func TestParseADTSProfile(t *testing.T) {
	header := []byte{0xFF, 0xF1, 0x40}
	prof := ParseADTSProfile(header)
	if prof != 1 {
		t.Errorf("expected profile 1, got %d", prof)
	}
}
