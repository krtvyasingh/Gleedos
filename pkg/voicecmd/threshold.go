package voicecmd

import "math"

func IsVoiceActive(samples []float64, thresholdRMS float64) bool {
	if len(samples) == 0 {
		return false
	}
	var sumSquares float64
	for _, s := range samples {
		sumSquares += s * s
	}
	rms := math.Sqrt(sumSquares / float64(len(samples)))
	return rms >= thresholdRMS
}
