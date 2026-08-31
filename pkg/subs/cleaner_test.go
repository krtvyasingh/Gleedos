package subs

import "testing"

func TestStripHTMLTags(t *testing.T) {
	clean := StripHTMLTags("<b>Hello</b> <i>World</i>")
	if clean != "Hello World" {
		t.Errorf("unexpected stripped text: %q", clean)
	}
}
