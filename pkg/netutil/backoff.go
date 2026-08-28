package netutil

import (
	"math/rand"
	"net/http"
	"time"
)

type BackoffTransport struct {
	Base       http.RoundTripper
	MaxRetries int
	BaseDelay  time.Duration
	MaxDelay   time.Duration
}

func NewBackoffTransport(base http.RoundTripper, maxRetries int) *BackoffTransport {
	if base == nil {
		base = http.DefaultTransport
	}
	if maxRetries <= 0 {
		maxRetries = 3
	}
	return &BackoffTransport{
		Base:       base,
		MaxRetries: maxRetries,
		BaseDelay:  100 * time.Millisecond,
		MaxDelay:   3 * time.Second,
	}
}

func (b *BackoffTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt <= b.MaxRetries; attempt++ {
		resp, err = b.Base.RoundTrip(req)
		if err == nil && resp.StatusCode != http.StatusTooManyRequests && resp.StatusCode != http.StatusServiceUnavailable && resp.StatusCode != http.StatusGatewayTimeout {
			return resp, nil
		}

		if attempt == b.MaxRetries {
			break
		}

		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}

		multiplier := time.Duration(1 << uint(attempt))
		delay := b.BaseDelay * multiplier
		if delay > b.MaxDelay {
			delay = b.MaxDelay
		}
		sleepDur := time.Duration(rand.Float64() * float64(delay))

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(sleepDur):
		}
	}
	return resp, err
}
