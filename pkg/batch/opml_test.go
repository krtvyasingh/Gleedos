package batch

import "testing"

func TestParseOPMLFeeds(t *testing.T) {
	opml := `<opml><body><outline xmlUrl="https://example.com/rss.xml"/><outline xmlUrl="https://podcast.com/feed"/></body></opml>`
	feeds := ParseOPMLFeeds(opml)
	if len(feeds) != 2 {
		t.Errorf("expected 2 feeds, got %d", len(feeds))
	}
}
