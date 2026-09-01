package tui

import "strings"

var WaveformBars = []rune{' ', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

func RenderWaveform(levels []float64) string {
	var sb strings.Builder
	for _, l := range levels {
		idx := int(l * float64(len(WaveformBars)-1))
		if idx < 0 {
			idx = 0
		}
		if idx >= len(WaveformBars) {
			idx = len(WaveformBars) - 1
		}
		sb.WriteRune(WaveformBars[idx])
	}
	return sb.String()
}
