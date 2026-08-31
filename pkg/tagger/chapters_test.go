package tagger

import (
	"testing"
	"time"
)

func TestChapterList(t *testing.T) {
	ch := []Chapter{{Title: "Intro", Start: 0, End: 10 * time.Second}}
	if FormatChapterList(ch) != 1 {
		t.Errorf("unexpected chapter count")
	}
}
