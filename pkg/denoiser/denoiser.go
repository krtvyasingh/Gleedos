package denoiser

import "math"

func ApplySpectralGate(spectrum []float64, noiseFloor float64) []float64 {
	out := make([]float64, len(spectrum))
	for i, mag := range spectrum {
		if mag < noiseFloor {
			out[i] = 0.0
		} else {
			out[i] = mag * (1.0 - math.Min(1.0, noiseFloor/(mag+1e-12)))
		}
	}
	return out
}
