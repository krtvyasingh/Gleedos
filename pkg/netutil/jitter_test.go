package netutil

import (
	"testing"
	"time"
)

func TestCalculateJitter(t *testing.T) {
	j1 := CalculateJitter(100*time.Millisecond, 1)
	j2 := CalculateJitter(100*time.Millisecond, 2)
	if j1 != 200*time.Millisecond || j2 != 400*time.Millisecond {
		t.Errorf("unexpected jitter: %v, %v", j1, j2)
	}
}
