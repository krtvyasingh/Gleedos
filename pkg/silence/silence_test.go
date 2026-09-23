package silence

import "testing"

func TestFindSilenceSegments(t *testing.T) {
	samples := make([]float64, 48000)
	for i := range samples {
		if i > 10000 && i < 20000 {
			samples[i] = 0.0
		} else {
			samples[i] = 0.5
		}
	}
	segs := FindSilenceSegments(samples, 0.01, 48000)
	if len(segs) == 0 {
		t.Errorf("expected silence detection")
	}
}
