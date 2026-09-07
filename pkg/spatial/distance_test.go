package spatial

import "testing"

func TestCalculateDistanceGain(t *testing.T) {
	g1 := CalculateDistanceGain(1.0, 1.0, 100.0)
	g2 := CalculateDistanceGain(2.0, 1.0, 100.0)
	if g1 != 1.0 || g2 != 0.5 {
		t.Errorf("unexpected gain: %f, %f", g1, g2)
	}
}
