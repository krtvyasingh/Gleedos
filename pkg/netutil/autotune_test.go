package netutil

import "testing"

func TestCalculateOptimalThreads(t *testing.T) {
	if CalculateOptimalThreads(5*1024*1024) != 2 {
		t.Errorf("expected 2 threads")
	}
	if CalculateOptimalThreads(500*1024*1024) != 8 {
		t.Errorf("expected 8 threads")
	}
}
