package loudness

import "math"

func CalculateTruePeak(samples []float64, oversample int) float64 {
	if len(samples) == 0 {
		return -70.0
	}
	var maxVal float64
	for _, s := range samples {
		absVal := math.Abs(s)
		if absVal > maxVal {
			maxVal = absVal
		}
	}
	if maxVal <= 1e-12 {
		return -70.0
	}
	return 20.0 * math.Log10(maxVal)
}
