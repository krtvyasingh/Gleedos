package subs

import (
	"testing"
	"time"
)

func TestIsValidCueDuration(t *testing.T) {
	if !IsValidCueDuration(1*time.Second, 4*time.Second) {
		t.Errorf("expected valid duration")
	}
	if IsValidCueDuration(4*time.Second, 1*time.Second) {
		t.Errorf("expected invalid duration")
	}
}
