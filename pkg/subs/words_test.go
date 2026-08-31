package subs

import "testing"

func TestParseWordCount(t *testing.T) {
	if ParseWordCount("Gleedos universal downloader engine") != 4 {
		t.Errorf("expected 4 words")
	}
}
