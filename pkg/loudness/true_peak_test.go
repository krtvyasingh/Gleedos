package loudness

import (
	"math"
	"testing"
)

func TestCalculateTruePeak(t *testing.T) {
	samples := []float64{0.1, -0.5, 0.8, -0.2}
	tp := CalculateTruePeak(samples, 4)
	expected := 20.0 * math.Log10(0.8)
	if math.Abs(tp-expected) > 0.01 {
		t.Errorf("unexpected true peak: %f (expected %f)", tp, expected)
	}
}
