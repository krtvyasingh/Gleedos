package tui

import "testing"

func TestRenderWaveform(t *testing.T) {
	wf := RenderWaveform([]float64{0.1, 0.5, 0.9})
	if len([]rune(wf)) != 3 {
		t.Errorf("unexpected waveform length: %s", wf)
	}
}
