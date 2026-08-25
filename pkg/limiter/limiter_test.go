package limiter

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"
)

func TestParseRate(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		hasErr   bool
	}{
		{"", 0, false},
		{"1024", 1024, false},
		{"1K", 1024, false},
		{"5M", 5 * 1024 * 1024, false},
		{"2.5MB", int64(2.5 * 1024 * 1024), false},
		{"1GB", 1024 * 1024 * 1024, false},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		got, err := ParseRate(tt.input)
		if (err != nil) != tt.hasErr {
			t.Errorf("ParseRate(%q) error = %v, expected error %v", tt.input, err, tt.hasErr)
		}
		if got != tt.expected {
			t.Errorf("ParseRate(%q) = %d, expected %d", tt.input, got, tt.expected)
		}
	}
}

func TestRateLimiter(t *testing.T) {
	// Limit to 10KB/s
	lim := New(10 * 1024)
	if lim == nil {
		t.Fatal("expected non-nil limiter")
	}

	data := make([]byte, 5*1024)
	reader := NewReader(context.Background(), bytes.NewReader(data), lim)

	start := time.Now()
	buf := make([]byte, 1024)
	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("unexpected read error: %v", err)
		}
	}
	elapsed := time.Since(start)
	t.Logf("Read 5KB took %v", elapsed)
}
