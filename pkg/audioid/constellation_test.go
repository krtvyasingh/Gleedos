package audioid

import "testing"

func TestExtractPeaks(t *testing.T) {
	spec := [][]float64{
		{0.1, 0.9, 0.2},
		{0.8, 0.0, 0.3},
	}
	peaks := ExtractPeaks(spec, 0.5)
	if len(peaks) != 2 {
		t.Errorf("expected 2 peaks, got %d", len(peaks))
	}
}
