package ui

import (
	"strings"
	"testing"
)

func TestFormatRow(t *testing.T) {
	row := FormatRow("Quality", "1080p")
	if !strings.Contains(row, "Quality") || !strings.Contains(row, "1080p") {
		t.Errorf("unexpected row: %s", row)
	}
}
