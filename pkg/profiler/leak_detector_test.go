package profiler

import "testing"

func TestLeakDetector(t *testing.T) {
	ld := NewLeakDetector()
	leaked, _ := ld.DetectDrift(1000)
	if leaked {
		t.Errorf("unexpected leak alert")
	}
}
