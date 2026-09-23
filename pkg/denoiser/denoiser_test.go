package denoiser

import "testing"

func TestApplySpectralGate(t *testing.T) {
	spec := []float64{0.05, 0.5, 0.02}
	clean := ApplySpectralGate(spec, 0.1)
	if clean[0] != 0.0 || clean[2] != 0.0 || clean[1] <= 0.0 {
		t.Errorf("unexpected denoised output: %+v", clean)
	}
}
