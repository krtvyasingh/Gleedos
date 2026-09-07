package tonemap

import "testing"

func TestReinhardExtended(t *testing.T) {
	val := ReinhardExtended(2.0, 4.0)
	if val < 0.6 || val > 0.9 {
		t.Errorf("unexpected tonemap result: %f", val)
	}
}
