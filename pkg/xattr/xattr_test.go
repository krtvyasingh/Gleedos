package xattr

import "testing"

func TestFormatXAttrMap(t *testing.T) {
	attrs := FormatXAttrMap(MediaMetadata{OriginalURL: "https://example.com", Title: "Sample"})
	if attrs["com.apple.metadata:kMDItemTitle"] != "Sample" {
		t.Errorf("unexpected xattr map: %+v", attrs)
	}
}
