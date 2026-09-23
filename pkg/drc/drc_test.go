package drc

import "testing"

func TestApplyDRC(t *testing.T) {
	s := ApplyDRC(1.0, 0.5, 2.0)
	if s != 0.75 {
		t.Errorf("expected 0.75, got %f", s)
	}
}
