package security

import "testing"

func TestIPRateLimiter(t *testing.T) {
	lim := NewIPRateLimiter(2)
	if !lim.Allow("127.0.0.1") || !lim.Allow("127.0.0.1") {
		t.Errorf("expected allowed")
	}
	if lim.Allow("127.0.0.1") {
		t.Errorf("expected rejected on 3rd request")
	}
}
