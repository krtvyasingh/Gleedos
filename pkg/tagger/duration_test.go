package tagger

import (
	"testing"
	"time"
)

func TestEstimateDuration(t *testing.T) {
	dur := EstimateDuration(160000, 128)
	if dur != 10*time.Second {
		t.Errorf("expected 10s, got %v", dur)
	}
}
