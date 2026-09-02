package profiler

import (
	"strings"
	"testing"
)

func TestCaptureStats(t *testing.T) {
	s := CaptureStats()
	f := FormatStats(s)
	if !strings.Contains(f, "Alloc:") {
		t.Errorf("unexpected format: %s", f)
	}
}
