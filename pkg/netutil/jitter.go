package netutil

import "time"

func CalculateJitter(base time.Duration, attempt int) time.Duration {
	mult := 1 << attempt
	if mult > 32 {
		mult = 32
	}
	return base * time.Duration(mult)
}
