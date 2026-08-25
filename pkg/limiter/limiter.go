package limiter

import (
	"context"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"
)

// RateLimiter enforces a maximum bytes-per-second bandwidth cap.
type RateLimiter struct {
	bytesPerSec int64
	tokens      float64
	capacity    float64
	lastRefill  time.Time
	mu          sync.Mutex
}

// New creates a RateLimiter with the given bytes per second. If bytesPerSec <= 0, no limiting is applied.
func New(bytesPerSec int64) *RateLimiter {
	if bytesPerSec <= 0 {
		return nil
	}
	cap := float64(bytesPerSec)
	return &RateLimiter{
		bytesPerSec: bytesPerSec,
		tokens:      cap,
		capacity:    cap,
		lastRefill:  time.Now(),
	}
}

// ParseRate parses human-readable rate strings like "5M", "500K", "10MB", "2048" into bytes per second.
func ParseRate(s string) (int64, error) {
	s = strings.TrimSpace(strings.ToUpper(s))
	if s == "" {
		return 0, nil
	}

	multiplier := int64(1)
	if strings.HasSuffix(s, "MB") || strings.HasSuffix(s, "M") {
		multiplier = 1024 * 1024
		s = strings.TrimSuffix(strings.TrimSuffix(s, "MB"), "M")
	} else if strings.HasSuffix(s, "KB") || strings.HasSuffix(s, "K") {
		multiplier = 1024
		s = strings.TrimSuffix(strings.TrimSuffix(s, "KB"), "K")
	} else if strings.HasSuffix(s, "GB") || strings.HasSuffix(s, "G") {
		multiplier = 1024 * 1024 * 1024
		s = strings.TrimSuffix(strings.TrimSuffix(s, "GB"), "G")
	} else if strings.HasSuffix(s, "B") {
		s = strings.TrimSuffix(s, "B")
	}

	val, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0, err
	}
	return int64(val * float64(multiplier)), nil
}

// Wait blocks until n bytes can be processed, or until ctx is cancelled.
func (r *RateLimiter) Wait(ctx context.Context, n int) error {
	if r == nil || r.bytesPerSec <= 0 || n <= 0 {
		return nil
	}

	for n > 0 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		r.mu.Lock()
		now := time.Now()
		elapsed := now.Sub(r.lastRefill).Seconds()
		r.lastRefill = now

		r.tokens += elapsed * float64(r.bytesPerSec)
		if r.tokens > r.capacity {
			r.tokens = r.capacity
		}

		needed := float64(n)
		if r.tokens >= needed {
			r.tokens -= needed
			r.mu.Unlock()
			return nil
		}

		// Take what's available
		available := r.tokens
		r.tokens = 0
		n -= int(available)
		sleepDuration := time.Duration((float64(n) / float64(r.bytesPerSec)) * float64(time.Second))
		if sleepDuration > 500*time.Millisecond {
			sleepDuration = 500 * time.Millisecond
		}
		if sleepDuration < 5*time.Millisecond {
			sleepDuration = 5 * time.Millisecond
		}
		r.mu.Unlock()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(sleepDuration):
		}
	}

	return nil
}

// Reader wraps an io.Reader with rate limiting.
type Reader struct {
	r   io.Reader
	lim *RateLimiter
	ctx context.Context
}

// NewReader returns an io.Reader limited by lim.
func NewReader(ctx context.Context, r io.Reader, lim *RateLimiter) io.Reader {
	if lim == nil {
		return r
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return &Reader{r: r, lim: lim, ctx: ctx}
}

func (lr *Reader) Read(p []byte) (int, error) {
	n, err := lr.r.Read(p)
	if n > 0 && lr.lim != nil {
		if waitErr := lr.lim.Wait(lr.ctx, n); waitErr != nil {
			return n, waitErr
		}
	}
	return n, err
}
