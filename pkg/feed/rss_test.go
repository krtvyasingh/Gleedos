package feed

import (
	"strings"
	"testing"
)

func TestParseRSS(t *testing.T) {
	xmlData := `<?xml version="1.0"?>
<rss version="2.0">
  <channel>
    <title>Podcast</title>
    <item>
      <title>Ep 1</title>
      <enclosure url="https://example.com/ep1.mp3"/>
    </item>
  </channel>
</rss>`

	feed, err := ParseRSS(strings.NewReader(xmlData))
	if err != nil || len(feed.Channel.Items) != 1 {
		t.Fatalf("ParseRSS failed: %v", err)
	}
	if feed.Channel.Items[0].Enclosure.URL != "https://example.com/ep1.mp3" {
		t.Errorf("unexpected enclosure URL")
	}
}
