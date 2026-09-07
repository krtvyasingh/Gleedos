package cdnroute

type LatencyEMA struct {
	alpha float64
	value float64
}

func NewLatencyEMA(alpha float64) *LatencyEMA {
	if alpha <= 0 || alpha > 1.0 {
		alpha = 0.2
	}
	return &LatencyEMA{alpha: alpha}
}

func (e *LatencyEMA) Update(sample float64) float64 {
	if e.value == 0 {
		e.value = sample
	} else {
		e.value = e.alpha*sample + (1.0-e.alpha)*e.value
	}
	return e.value
}
