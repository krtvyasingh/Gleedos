package audioid

import "testing"

func TestNormalizeBin(t *testing.T) {
	v := NormalizeBin(5.0, 0.0, 10.0)
	if v != 0.5 {
		t.Errorf("expected 0.5, got %f", v)
	}
}
