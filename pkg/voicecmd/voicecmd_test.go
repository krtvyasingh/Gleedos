package voicecmd

import "testing"

func TestMatchHotword(t *testing.T) {
	matched, cmd := MatchHotword("Hey gleedos download https://example.com/video")
	if !matched || cmd != "https://example.com/video" {
		t.Errorf("hotword matching failed: %v, %s", matched, cmd)
	}
}
