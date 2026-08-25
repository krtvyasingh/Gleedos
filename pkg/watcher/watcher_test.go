package watcher

import (
	"testing"
)

func TestIsMediaURL(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"https://www.youtube.com/watch?v=dQw4w9WgXcQ", true},
		{"https://youtu.be/dQw4w9WgXcQ", true},
		{"https://vimeo.com/12345678", true},
		{"https://example.com/video.mp4", true},
		{"https://example.com/audio.mp3?token=abc", true},
		{"https://example.com/stream.m3u8", true},
		{"https://google.com/search?q=golang", false},
		{"random string text", false},
	}

	for _, tt := range tests {
		got := IsMediaURL(tt.url)
		if got != tt.expected {
			t.Errorf("IsMediaURL(%q) = %v, expected %v", tt.url, got, tt.expected)
		}
	}
}
