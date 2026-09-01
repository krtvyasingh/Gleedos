package loudness

import "testing"

func TestCalculateLUFS(t *testing.T) {
	samples := make([]float64, 1000)
	for i := range samples {
		samples[i] = 0.5
	}
	res := CalculateLUFS(samples)
	if res.IntegratedLUFS < -20.0 || res.IntegratedLUFS > 0.0 {
		t.Errorf("unexpected LUFS: %+v", res)
	}
}
