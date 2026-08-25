package ui

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// ProgressTracker tracks and renders live download stats.
type ProgressTracker struct {
	TotalBytes   int64
	CurrentBytes int64
	StartTime    time.Time
	LastUpdate   time.Time
	LastBytes    int64
	Speed        float64 // bytes/sec
	WorkerCount  int
	Status       string
	mu           sync.Mutex
	writer       io.Writer
	quiet        bool
}

// NewProgressTracker creates a progress tracker.
func NewProgressTracker(totalBytes int64, workerCount int) *ProgressTracker {
	return &ProgressTracker{
		TotalBytes:  totalBytes,
		WorkerCount: workerCount,
		StartTime:   time.Now(),
		LastUpdate:  time.Now(),
		writer:      os.Stdout,
		Status:      "Downloading",
	}
}

// SetQuiet disables progress rendering.
func (p *ProgressTracker) SetQuiet(quiet bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.quiet = quiet
}

// SetWriter sets custom output writer (useful for tests).
func (p *ProgressTracker) SetWriter(w io.Writer) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.writer = w
}

// AddBytes increments current transferred byte count and recalculates speed.
func (p *ProgressTracker) AddBytes(n int64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.CurrentBytes += n
	now := time.Now()
	dur := now.Sub(p.LastUpdate).Seconds()
	if dur >= 0.25 {
		bytesDelta := p.CurrentBytes - p.LastBytes
		instantSpeed := float64(bytesDelta) / dur
		if p.Speed == 0 {
			p.Speed = instantSpeed
		} else {
			p.Speed = 0.7*p.Speed + 0.3*instantSpeed // EMA smoothing
		}
		p.LastBytes = p.CurrentBytes
		p.LastUpdate = now
	}
}

// SetStatus updates current status label.
func (p *ProgressTracker) SetStatus(status string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Status = status
}

// FormatBytes converts raw bytes into human readable format (KB, MB, GB).
func FormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

// Render returns a formatted single-line status string.
func (p *ProgressTracker) Render() string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.TotalBytes <= 0 {
		speedStr := FormatBytes(int64(p.Speed)) + "/s"
		return fmt.Sprintf("\r  [%s] %s | Speed: %s | Workers: %d",
			p.Status, FormatBytes(p.CurrentBytes), speedStr, p.WorkerCount)
	}

	percent := float64(p.CurrentBytes) / float64(p.TotalBytes) * 100.0
	if percent > 100.0 {
		percent = 100.0
	}

	// Bar width: 25 characters
	barWidth := 25
	filled := int((percent / 100.0) * float64(barWidth))
	if filled > barWidth {
		filled = barWidth
	}
	empty := barWidth - filled
	bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)

	speedStr := FormatBytes(int64(p.Speed)) + "/s"

	var etaStr string
	if p.Speed > 0 && p.CurrentBytes < p.TotalBytes {
		remainingSec := float64(p.TotalBytes-p.CurrentBytes) / p.Speed
		eta := time.Duration(remainingSec) * time.Second
		etaStr = fmt.Sprintf("ETA: %02d:%02d", int(eta.Minutes()), int(eta.Seconds())%60)
	} else if p.CurrentBytes >= p.TotalBytes {
		etaStr = "Done"
	} else {
		etaStr = "ETA: --:--"
	}

	return fmt.Sprintf("\r  [%s] [%s] %5.1f%% | %s/%s | %s | %s | W:%d",
		p.Status, bar, percent, FormatBytes(p.CurrentBytes), FormatBytes(p.TotalBytes), speedStr, etaStr, p.WorkerCount)
}

// PrintProgress writes the rendered progress line to stdout.
func (p *ProgressTracker) PrintProgress() {
	if p.quiet {
		return
	}
	line := p.Render()
	p.mu.Lock()
	w := p.writer
	p.mu.Unlock()
	fmt.Fprint(w, line)
}

// Finish prints completion status and moves to the next terminal line.
func (p *ProgressTracker) Finish() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.quiet {
		return
	}
	fmt.Fprintln(p.writer)
}
