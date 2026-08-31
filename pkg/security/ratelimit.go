package security

import "sync"

type IPRateLimiter struct {
	counts map[string]int
	mu     sync.Mutex
	limit  int
}

func NewIPRateLimiter(limit int) *IPRateLimiter {
	return &IPRateLimiter{counts: make(map[string]int), limit: limit}
}

func (l *IPRateLimiter) Allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.counts[ip] >= l.limit {
		return false
	}
	l.counts[ip]++
	return true
}
