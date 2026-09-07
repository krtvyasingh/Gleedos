package airplay

import "testing"

func TestParseAirPlayCommand(t *testing.T) {
	cmd := ParseAirPlayCommand("POST /play HTTP/1.1\r\nContent-Length: 0\r\n\r\n")
	if cmd != "play" {
		t.Errorf("expected play command, got %s", cmd)
	}
}
