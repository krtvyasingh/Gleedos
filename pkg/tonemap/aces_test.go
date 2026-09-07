package tonemap

import "testing"

func TestACESFilmicTonemap(t *testing.T) {
	v := ACESFilmicTonemap(1.0)
	if v < 0.7 || v > 0.9 {
		t.Errorf("unexpected ACES value: %f", v)
	}
}
