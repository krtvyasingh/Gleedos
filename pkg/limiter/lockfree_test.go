package limiter

import "testing"

func TestLockFreeBucket(t *testing.T) {
	b := NewLockFreeBucket(10)
	if !b.TryConsume(5) {
		t.Errorf("expected consume success")
	}
	if b.TryConsume(10) {
		t.Errorf("expected consume failure")
	}
}
