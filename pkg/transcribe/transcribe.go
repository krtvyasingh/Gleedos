package transcribe

import (
	"fmt"
	"io"
	"time"
)

type SubtitleCue struct {
	Index int
	Start time.Duration
	End   time.Duration
	Text  string
}

func FormatSRTCue(cue SubtitleCue) string {
	startFmt := formatTime(cue.Start, ",")
	endFmt := formatTime(cue.End, ",")
	return fmt.Sprintf("%d\n%s --> %s\n%s\n", cue.Index, startFmt, endFmt, cue.Text)
}

func formatTime(d time.Duration, delim string) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	ms := d / time.Millisecond
	return fmt.Sprintf("%02d:%02d:%02d%s%03d", h, m, s, delim, ms)
}

func WriteSRT(cues []SubtitleCue, w io.Writer) error {
	for _, c := range cues {
		if _, err := io.WriteString(w, FormatSRTCue(c)+"\n"); err != nil {
			return err
		}
	}
	return nil
}
