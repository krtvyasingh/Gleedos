package torrent

import "testing"

func TestNewWebSeed(t *testing.T) {
	ws := NewWebSeed("https://example.com/file.mp4")
	if ws.URL != "https://example.com/file.mp4" {
		t.Errorf("unexpected webseed url")
	}
}
