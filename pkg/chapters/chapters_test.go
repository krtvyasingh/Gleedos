package chapters

import (
	"strings"
	"testing"
	"time"
)

func TestFormatFFMetadata(t *testing.T) {
	ch := []VideoChapter{
		{Title: "Intro", Timestamp: 0},
		{Title: "Main Topic", Timestamp: 30 * time.Second},
	}
	meta := FormatFFMetadata(ch)
	if !strings.Contains(meta, "title=Intro") || !strings.Contains(meta, "title=Main Topic") {
		t.Errorf("unexpected metadata: %s", meta)
	}
}
