package ui

import "testing"

func TestGetSpinnerFrame(t *testing.T) {
	frame := GetSpinnerFrame(0)
	if frame != "⠋" {
		t.Errorf("unexpected frame: %s", frame)
	}
}
