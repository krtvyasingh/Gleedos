package watcher

import (
	"context"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ReadClipboard returns current clipboard text using native OS tools.
func ReadClipboard() (string, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbpaste")
	case "windows":
		cmd = exec.Command("powershell.exe", "-NoProfile", "-Command", "Get-Clipboard")
	default:
		// Linux: try xclip then xsel
		if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard", "-o")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "-b", "-o")
		} else {
			return "", nil
		}
	}

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// IsMediaURL checks if a string is a recognized media URL.
func IsMediaURL(s string) bool {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
		return false
	}

	lower := strings.ToLower(s)
	mediaDomains := []string{
		"youtube.com", "youtu.be",
		"vimeo.com", "twitch.tv",
		"twitter.com", "x.com",
		"tiktok.com", "instagram.com",
		"soundcloud.com", "reddit.com",
		"facebook.com", "fb.watch",
	}

	for _, d := range mediaDomains {
		if strings.Contains(lower, d) {
			return true
		}
	}

	mediaExts := []string{
		".mp4", ".mkv", ".webm", ".m3u8",
		".mp3", ".flac", ".aac", ".wav",
		".mov", ".avi", ".ts",
	}

	for _, ext := range mediaExts {
		if strings.HasSuffix(lower, ext) || strings.Contains(lower, ext+"?") {
			return true
		}
	}

	return false
}

// Watcher monitors clipboard for new media URLs.
type Watcher struct {
	PollInterval time.Duration
	Handler      func(url string)
	seen         map[string]time.Time
	mu           sync.Mutex
}

// New creates a new clipboard Watcher.
func New(pollInterval time.Duration, handler func(url string)) *Watcher {
	if pollInterval <= 0 {
		pollInterval = 1 * time.Second
	}
	return &Watcher{
		PollInterval: pollInterval,
		Handler:      handler,
		seen:         make(map[string]time.Time),
	}
}

// Start begins polling clipboard until context is cancelled.
func (w *Watcher) Start(ctx context.Context) {
	ticker := time.NewTicker(w.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			text, err := ReadClipboard()
			if err != nil || text == "" {
				continue
			}

			if !IsMediaURL(text) {
				continue
			}

			w.mu.Lock()
			lastSeen, exists := w.seen[text]
			// Avoid re-triggering if seen within last 10 minutes
			if exists && time.Since(lastSeen) < 10*time.Minute {
				w.mu.Unlock()
				continue
			}
			w.seen[text] = time.Now()
			w.mu.Unlock()

			if w.Handler != nil {
				go w.Handler(text)
			}
		}
	}
}
