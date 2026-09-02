package tonemap

import "testing"

func TestReinhardTonemap(t *testing.T) {
	val := ReinhardTonemap(2.0)
	if val < 0.66 || val > 0.67 {
		t.Errorf("unexpected tonemapped value: %f", val)
	}
}
