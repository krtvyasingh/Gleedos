package urlutil

import "testing"

func TestDeduplicateURLs(t *testing.T) {
	raw := []string{
		"https://example.com/video?utm_source=1",
		"https://example.com/video?utm_source=2",
		"https://example.com/other",
	}

	deduped := DeduplicateURLs(raw)
	if len(deduped) != 2 {
		t.Errorf("expected 2 unique URLs, got %d", len(deduped))
	}
}
