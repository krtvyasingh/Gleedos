package notion

import (
	"strings"
	"testing"
)

func TestFormatCalloutBlock(t *testing.T) {
	callout := FormatCalloutBlock("💡", "Pro Tip")
	if !strings.Contains(callout, "💡 Pro Tip") {
		t.Errorf("unexpected callout: %s", callout)
	}
}
