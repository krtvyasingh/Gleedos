package notion

import (
	"strings"
	"testing"
)

func TestFormatNotionPageJSON(t *testing.T) {
	json := FormatNotionPageJSON("My Video", "https://example.com")
	if !strings.Contains(json, "My Video") || !strings.Contains(json, "https://example.com") {
		t.Errorf("unexpected JSON: %s", json)
	}
}
