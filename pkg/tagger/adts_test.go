package tagger

import "testing"

func TestIsADTS(t *testing.T) {
	if !IsADTS([]byte{0xFF, 0xF1, 0x00}) {
		t.Errorf("expected true for ADTS syncword")
	}
}
