package subs

import (
	"bytes"
	"strings"
	"testing"
)

func TestVTTtoSRT(t *testing.T) {
	vtt := `WEBVTT

00:00:01.000 --> 00:00:04.000
Hello World!
`
	var srt bytes.Buffer
	if err := VTTtoSRT(strings.NewReader(vtt), &srt); err != nil {
		t.Fatalf("VTTtoSRT failed: %v", err)
	}

	out := srt.String()
	if !strings.Contains(out, "00:00:01,000 --> 00:00:04,000") || !strings.Contains(out, "Hello World!") {
		t.Errorf("unexpected SRT output: %s", out)
	}
}
