package demux

import "testing"

func TestIsWebM(t *testing.T) {
	hdr := append([]byte{0x1A, 0x45, 0xDF, 0xA3}, []byte("...webm...")...)
	if !IsWebM(hdr) {
		t.Errorf("expected true for WebM header")
	}
}
