package dht

import "testing"

func TestFormatPeerCompact(t *testing.T) {
	p := FormatPeerCompact(PeerEndpoint{IP: "1.2.3.4", Port: 6881})
	if p != "1.2.3.4" {
		t.Errorf("unexpected format: %s", p)
	}
}
