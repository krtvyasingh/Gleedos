package demux

import "testing"

func TestIsMKV(t *testing.T) {
	if !IsMKV([]byte{0x1A, 0x45, 0xDF, 0xA3, 0x01}) {
		t.Errorf("expected true for MKV EBML header")
	}
}
