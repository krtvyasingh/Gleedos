package ui

import (
	"strings"
	"testing"
)

func TestRenderSlot(t *testing.T) {
	str := RenderSlot(BarSlot{ID: 1, Percent: 50.0, Label: "Video"})
	if !strings.Contains(str, "[1] Video") || !strings.Contains(str, "50.0%") {
		t.Errorf("unexpected slot output: %s", str)
	}
}
