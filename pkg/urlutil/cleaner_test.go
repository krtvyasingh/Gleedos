package urlutil

import "testing"

func TestCleanURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://example.com/video?utm_source=twitter&v=123", "https://example.com/video?v=123"},
		{"https://youtu.be/abc?si=xyz&t=40", "https://youtu.be/abc?t=40"},
		{"https://example.com/clean", "https://example.com/clean"},
	}

	for _, tt := range tests {
		got := CleanURL(tt.input)
		if got != tt.expected {
			t.Errorf("CleanURL(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}
