package subs

import "testing"

func TestIsValidUTF8(t *testing.T) {
	if !IsValidUTF8([]byte("Hello 世界")) {
		t.Errorf("expected valid UTF-8")
	}
}
