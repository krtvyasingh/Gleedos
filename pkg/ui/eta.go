package ui

import "time"

type ETACalculator struct {
	smoothedSpeed float64
}

func (e *ETACalculator) Update(instantSpeed float64) {
	if e.smoothedSpeed == 0 {
		e.smoothedSpeed = instantSpeed
	} else {
		e.smoothedSpeed = 0.8*e.smoothedSpeed + 0.2*instantSpeed
	}
}

func (e *ETACalculator) EstimateRemaining(remainingBytes int64) time.Duration {
	if e.smoothedSpeed <= 0 {
		return 0
	}
	return time.Duration(float64(remainingBytes)/e.smoothedSpeed) * time.Second
}
