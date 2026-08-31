package subs

import (
	"testing"
	"time"
)

func TestShiftTimestamp(t *testing.T) {
	res := ShiftTimestamp(5*time.Second, 2*time.Second)
	if res != 7*time.Second {
		t.Errorf("expected 7s, got %v", res)
	}
}
