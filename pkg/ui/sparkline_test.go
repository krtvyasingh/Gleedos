package ui

import (
	"testing"
)

func TestRenderSparkline(t *testing.T) {
	vals := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 8.0}
	spark := RenderSparkline(vals)
	if len([]rune(spark)) != len(vals) {
		t.Errorf("expected sparkline length %d, got %d", len(vals), len([]rune(spark)))
	}
}
