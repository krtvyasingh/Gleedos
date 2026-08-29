package ui

import "strings"

var sparkTicks = []rune{' ', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

func RenderSparkline(values []float64) string {
	if len(values) == 0 {
		return ""
	}

	min, max := values[0], values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	var sb strings.Builder
	for _, v := range values {
		idx := 0
		if max > min {
			ratio := (v - min) / (max - min)
			idx = int(ratio * float64(len(sparkTicks)-1))
			if idx >= len(sparkTicks) {
				idx = len(sparkTicks) - 1
			}
		}
		sb.WriteRune(sparkTicks[idx])
	}
	return sb.String()
}
