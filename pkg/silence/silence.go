package silence

import "math"

func FindSilenceSegments(samples []float64, thresholdRMS float64, sampleRate int) [][2]int {
	var intervals [][2]int
	inSilence := false
	start := 0
	window := sampleRate / 10 // 100ms
	if window <= 0 {
		window = 1
	}

	for i := 0; i < len(samples); i += window {
		end := i + window
		if end > len(samples) {
			end = len(samples)
		}
		var sumSquares float64
		for _, s := range samples[i:end] {
			sumSquares += s * s
		}
		rms := math.Sqrt(sumSquares / float64(end-i))
		if rms < thresholdRMS {
			if !inSilence {
				inSilence = true
				start = i
			}
		} else {
			if inSilence {
				inSilence = false
				intervals = append(intervals, [2]int{start, i})
			}
		}
	}
	return intervals
}
