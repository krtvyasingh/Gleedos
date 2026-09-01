package transcribe

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestWriteSRT(t *testing.T) {
	cues := []SubtitleCue{
		{Index: 1, Start: 1 * time.Second, End: 4 * time.Second, Text: "Hello Whisper"},
	}
	var buf bytes.Buffer
	if err := WriteSRT(cues, &buf); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "00:00:01,000 --> 00:00:04,000") {
		t.Errorf("unexpected SRT output: %s", buf.String())
	}
}
