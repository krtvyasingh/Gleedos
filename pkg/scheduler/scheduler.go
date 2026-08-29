package scheduler

import (
	"time"
)

type TimeWindow struct {
	StartHour int
	StartMin  int
	EndHour   int
	EndMin    int
}

func (tw TimeWindow) IsActive(t time.Time) bool {
	hour, min, _ := t.Clock()
	cur := hour*60 + min
	start := tw.StartHour*60 + tw.StartMin
	end := tw.EndHour*60 + tw.EndMin

	if start <= end {
		return cur >= start && cur <= end
	}
	return cur >= start || cur <= end
}
