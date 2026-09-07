package safety

import "testing"

func TestSafetyFilter(t *testing.T) {
	sf := NewSafetyFilter([]string{"danger", "toxic"})
	if !sf.ContainsBanned("This is a toxic sample") {
		t.Errorf("expected banned word detection")
	}
	if sf.ContainsBanned("Clean and safe content") {
		t.Errorf("expected safe pass")
	}
}
