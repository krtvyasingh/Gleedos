package ui

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		bytes    int64
		expected string
	}{
		{500, "500 B"},
		{1024, "1.00 KB"},
		{1048576, "1.00 MB"},
		{1073741824, "1.00 GB"},
	}

	for _, tt := range tests {
		got := FormatBytes(tt.bytes)
		if got != tt.expected {
			t.Errorf("FormatBytes(%d) = %q, expected %q", tt.bytes, got, tt.expected)
		}
	}
}

func TestProgressTracker(t *testing.T) {
	var buf bytes.Buffer
	tracker := NewProgressTracker(10000, 4)
	tracker.SetWriter(&buf)

	tracker.AddBytes(5000)
	time.Sleep(10 * time.Millisecond)
	tracker.PrintProgress()

	rendered := buf.String()
	if !strings.Contains(rendered, "50.0%") && !strings.Contains(rendered, "Downloading") {
		t.Errorf("rendered line unexpected: %s", rendered)
	}

	tracker.Finish()
}
