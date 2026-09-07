package visionllm

import "testing"

func TestIsSceneChange(t *testing.T) {
	h1 := []float64{0.1, 0.2, 0.3}
	h2 := []float64{0.9, 0.8, 0.7}
	if !IsSceneChange(h1, h2, 0.5) {
		t.Errorf("expected scene change")
	}
	if IsSceneChange(h1, h1, 0.5) {
		t.Errorf("expected no scene change for identical histograms")
	}
}
