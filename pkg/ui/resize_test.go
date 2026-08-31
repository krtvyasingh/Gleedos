package ui

import "testing"

func TestDefaultTerminalWidth(t *testing.T) {
	if DefaultTerminalWidth() != 80 {
		t.Errorf("expected 80")
	}
}
