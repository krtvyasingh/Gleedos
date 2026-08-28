package hls

import "testing"

func TestParseByteRange(t *testing.T) {
	br, err := ParseByteRange("#EXT-X-BYTERANGE:1024@2048", 0)
	if err != nil || br.Length != 1024 || br.Offset != 2048 {
		t.Errorf("ParseByteRange failed: %v, %+v", err, br)
	}
}
