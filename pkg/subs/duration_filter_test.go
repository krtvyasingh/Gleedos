package subs

import (
	"testing"
	"time"
)

func TestExtendShortCues(t *testing.T) {
	d1 := ExtendShortCues(500*time.Millisecond, 1*time.Second)
	d2 := ExtendShortCues(2*time.Second, 1*time.Second)
	if d1 != 1*time.Second || d2 != 2*time.Second {
		t.Errorf("unexpected cue durations: %v, %v", d1, d2)
	}
}
