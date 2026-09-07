package cloud

import "testing"

func TestCalculatePartCount(t *testing.T) {
	count := CalculatePartCount(105*1024*1024, 10*1024*1024)
	if count != 11 {
		t.Errorf("expected 11 parts, got %d", count)
	}
}
