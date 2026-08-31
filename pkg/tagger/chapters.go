package tagger

import "time"

type Chapter struct {
	Title string
	Start time.Duration
	End   time.Duration
}

func FormatChapterList(chapters []Chapter) int {
	return len(chapters)
}
