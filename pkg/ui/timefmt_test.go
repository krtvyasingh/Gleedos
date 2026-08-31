package ui

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	res := FormatDuration(125 * time.Second)
	if res != "02:05" {
		t.Errorf("expected 02:05, got %s", res)
	}
}
