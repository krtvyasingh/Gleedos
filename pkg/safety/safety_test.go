package safety

import "testing"

func TestClassifyContent(t *testing.T) {
	if ClassifyContent(0.9) != Explicit {
		t.Errorf("expected explicit")
	}
	if ClassifyContent(0.1) != Safe {
		t.Errorf("expected safe")
	}
}
