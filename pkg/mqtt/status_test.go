package mqtt

import (
	"strings"
	"testing"
)

func TestFormatClientStatusJSON(t *testing.T) {
	s := FormatClientStatusJSON("node-1", "online")
	if !strings.Contains(s, `"state":"online"`) {
		t.Errorf("unexpected status JSON: %s", s)
	}
}
