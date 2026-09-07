package visionllm

import "time"

func CalculateKeyframeInterval(duration time.Duration, maxFrames int) time.Duration {
	if maxFrames <= 0 {
		return time.Second
	}
	interval := duration / time.Duration(maxFrames)
	if interval < 500*time.Millisecond {
		return 500 * time.Millisecond
	}
	return interval
}
