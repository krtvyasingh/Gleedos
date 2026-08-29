package subs

import (
	"bytes"
	"strings"
	"testing"
)

func TestSRTtoASS(t *testing.T) {
	srt := `1
00:00:01,000 --> 00:00:04,000
First subtitle line

2
00:00:05,500 --> 00:00:08,000
Second subtitle line
`
	var ass bytes.Buffer
	if err := SRTtoASS(strings.NewReader(srt), &ass); err != nil {
		t.Fatalf("SRTtoASS failed: %v", err)
	}

	out := ass.String()
	if !strings.Contains(out, "[V4+ Styles]") || !strings.Contains(out, "Dialogue: 0,00:00:01.00,00:00:04.00,Default") {
		t.Errorf("unexpected ASS output: %s", out)
	}
}
