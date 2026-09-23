package drc

import "math"

func ApplyDRC(sample float64, threshold, ratio float64) float64 {
	absVal := math.Abs(sample)
	if absVal <= threshold {
		return sample
	}
	excess := absVal - threshold
	compressed := threshold + (excess / ratio)
	if sample < 0 {
		return -compressed
	}
	return compressed
}
