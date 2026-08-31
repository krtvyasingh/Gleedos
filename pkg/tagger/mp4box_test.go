package tagger

import "testing"

func TestNewBox(t *testing.T) {
	b := NewBox("moov")
	if b.Name != "moov" || len(b.Children) != 0 {
		t.Errorf("unexpected box: %+v", b)
	}
}
