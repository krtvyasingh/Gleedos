package chapters

import (
	"fmt"
	"time"
)

type VideoChapter struct {
	Title     string
	Timestamp time.Duration
}

func FormatFFMetadata(chapters []VideoChapter) string {
	out := ";FFMETADATA1\n"
	for i, c := range chapters {
		var endTime time.Duration
		if i+1 < len(chapters) {
			endTime = chapters[i+1].Timestamp
		} else {
			endTime = c.Timestamp + 60*time.Second
		}
		out += fmt.Sprintf("[CHAPTER]\nTIMEBASE=1/1000\nSTART=%d\nEND=%d\ntitle=%s\n",
			c.Timestamp.Milliseconds(), endTime.Milliseconds(), c.Title)
	}
	return out
}
