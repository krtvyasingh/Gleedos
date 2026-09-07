package tcptune

import "testing"

func TestCalculateOptimalWindow(t *testing.T) {
	win := CalculateOptimalWindow(100*1024*1024, 20)
	if win < 64*1024 || win > 32*1024*1024 {
		t.Errorf("unexpected window: %d", win)
	}
}
