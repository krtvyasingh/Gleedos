package audioid

import "testing"

func TestFingerprintSamples(t *testing.T) {
	samples := []float64{0.1, 0.2, 0.3}
	fp := FingerprintSamples(samples)
	if len(fp) != 64 {
		t.Errorf("invalid fingerprint: %s", fp)
	}
}
