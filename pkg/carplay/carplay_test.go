package carplay

import (
	"strings"
	"testing"
)

func TestGenerateCarPlayFeed(t *testing.T) {
	feed := GenerateCarPlayFeed("My Car Playlist", "http://192.168.1.50/audio.mp3")
	if !strings.Contains(feed, "My Car Playlist") || !strings.Contains(feed, "audio/mpeg") {
		t.Errorf("unexpected RSS feed: %s", feed)
	}
}
