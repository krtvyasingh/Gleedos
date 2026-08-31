package subs

import "time"

func IsValidCueDuration(start, end time.Duration) bool {
	return end > start && (end-start) < 30*time.Second
}
