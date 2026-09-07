package demux

import "testing"

func TestParseFLACBlockHeader(t *testing.T) {
	buf := []byte{0x84, 0x00, 0x01, 0x00}
	h := ParseFLACBlockHeader(buf)
	if !h.IsLast || h.BlockType != 4 || h.Length != 256 {
		t.Errorf("unexpected FLAC header: %+v", h)
	}
}
