package netutil

import "testing"

func TestCongestionController(t *testing.T) {
	cc := NewCongestionController()
	cc.OnSuccess()
	if cc.WindowSize != 5 {
		t.Errorf("expected window size 5, got %d", cc.WindowSize)
	}
	cc.OnDrop()
	if cc.WindowSize != 2 {
		t.Errorf("expected window size 2, got %d", cc.WindowSize)
	}
}
