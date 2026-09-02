package scraper

import "testing"

func TestExtractMediaURLs(t *testing.T) {
	html := `<div><a href="https://example.com/media.m3u8">Stream</a><video src="https://example.com/vid.mp4"></video></div>`
	urls := ExtractMediaURLs(html)
	if len(urls) != 2 {
		t.Errorf("expected 2 URLs, got %d", len(urls))
	}
}
