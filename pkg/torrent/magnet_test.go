package torrent

import "testing"

func TestParseMagnet(t *testing.T) {
	m, err := ParseMagnet("magnet:?xt=urn:btih:1234567890abcdef&dn=Video+Title")
	if err != nil || m.DisplayName != "Video Title" {
		t.Fatalf("ParseMagnet failed: %v, %+v", err, m)
	}
}
