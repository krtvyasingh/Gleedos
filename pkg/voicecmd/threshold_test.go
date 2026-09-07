package voicecmd

import "testing"

func TestIsVoiceActive(t *testing.T) {
	samples := []float64{0.5, 0.5, 0.5}
	if !IsVoiceActive(samples, 0.1) {
		t.Errorf("expected active voice")
	}
	if IsVoiceActive([]float64{0.001, 0.001}, 0.1) {
		t.Errorf("expected inactive voice")
	}
}
