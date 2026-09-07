package visionllm

import "math"

func IsSceneChange(histA, histB []float64, threshold float64) bool {
	if len(histA) != len(histB) || len(histA) == 0 {
		return false
	}
	var diff float64
	for i := range histA {
		diff += math.Abs(histA[i] - histB[i])
	}
	return (diff / float64(len(histA))) > threshold
}
