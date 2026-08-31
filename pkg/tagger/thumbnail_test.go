package tagger

import "testing"

func TestNewThumbnail(t *testing.T) {
	th := NewThumbnail("image/jpeg", []byte{0xFF, 0xD8, 0xFF})
	if th.MimeType != "image/jpeg" || len(th.Data) != 3 {
		t.Errorf("unexpected thumbnail: %+v", th)
	}
}
