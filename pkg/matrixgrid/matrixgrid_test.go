package matrixgrid

import (
	"strings"
	"testing"
)

func TestRenderSlotStatus(t *testing.T) {
	s := RenderSlotStatus(1, "Video", 85.5)
	if !strings.Contains(s, "[Slot 1] Video: 85.5%") {
		t.Errorf("unexpected slot format: %s", s)
	}
}
