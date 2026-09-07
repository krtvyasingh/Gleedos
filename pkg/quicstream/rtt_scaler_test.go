package quicstream

import (
	"testing"
	"time"
)

func TestScaleCongestionWindow(t *testing.T) {
	w1 := ScaleCongestionWindow(5*time.Millisecond, 100)
	w2 := ScaleCongestionWindow(300*time.Millisecond, 100)
	if w1 != 200 || w2 != 50 {
		t.Errorf("unexpected scaled windows: %d, %d", w1, w2)
	}
}
