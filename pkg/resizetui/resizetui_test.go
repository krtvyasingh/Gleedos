package resizetui

import "testing"

func TestClampDimensions(t *testing.T) {
	d := ClampDimensions(20, 5)
	if d.Width != 40 || d.Height != 10 {
		t.Errorf("unexpected clamped dimensions: %+v", d)
	}
}
