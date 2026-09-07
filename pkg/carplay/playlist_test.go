package carplay

import (
	"strings"
	"testing"
)

func TestFormatPlaylistItem(t *testing.T) {
	item := FormatPlaylistItem("Track", "http://media/1.mp3", "03:30")
	if !strings.Contains(item, "<title>Track</title>") {
		t.Errorf("unexpected playlist item: %s", item)
	}
}
