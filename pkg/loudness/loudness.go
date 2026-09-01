package loudness

import "math"

type LoudnessResult struct {
	IntegratedLUFS float64
	PeakLUFS       float64
	GainAdjustment float64
}

func CalculateLUFS(samples []float64) LoudnessResult {
	if len(samples) == 0 {
		return LoudnessResult{IntegratedLUFS: -70.0, PeakLUFS: -70.0, GainAdjustment: 0.0}
	}
	var sumSquares float64
	var peak float64
	for _, s := range samples {
		absS := math.Abs(s)
		if absS > peak {
			peak = absS
		}
		sumSquares += s * s
	}
	meanSquare := sumSquares / float64(len(samples))
	lufs := -0.691 + 10*math.Log10(math.Max(meanSquare, 1e-12))
	targetLUFS := -14.0 // YouTube / Streaming standard
	return LoudnessResult{
		IntegratedLUFS: lufs,
		PeakLUFS:       20 * math.Log10(math.Max(peak, 1e-12)),
		GainAdjustment: targetLUFS - lufs,
	}
}
