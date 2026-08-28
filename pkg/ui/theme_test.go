package ui

import (
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	res := Colorize(ColorGreen, "Success")
	if !strings.Contains(res, "Success") || !strings.HasPrefix(res, "[32m") {
		t.Errorf("unexpected colored string: %q", res)
	}
}
