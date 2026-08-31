package ui

import (
	"testing"
	"time"
)

func TestETACalculator(t *testing.T) {
	eta := &ETACalculator{}
	eta.Update(1000)
	rem := eta.EstimateRemaining(10000)
	if rem != 10*time.Second {
		t.Errorf("expected 10s, got %v", rem)
	}
}
