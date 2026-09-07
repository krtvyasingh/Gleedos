package subs

import "time"

func ExtendShortCues(duration, minDuration time.Duration) time.Duration {
	if duration < minDuration {
		return minDuration
	}
	return duration
}
