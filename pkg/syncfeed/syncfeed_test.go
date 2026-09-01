package syncfeed

import "testing"

func TestHasNewItem(t *testing.T) {
	fc := &FeedChannel{FeedURL: "https://example.com/rss", LastItem: "item1"}
	if !HasNewItem(fc, "item2") {
		t.Errorf("expected new item detected")
	}
}
