package subs

import "time"

func ShiftTimestamp(t time.Duration, delta time.Duration) time.Duration {
	res := t + delta
	if res < 0 {
		return 0
	}
	return res
}
