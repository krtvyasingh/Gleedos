package quicstream

import "time"

func ScaleCongestionWindow(rtt time.Duration, baseWindow int) int {
	if rtt < 10*time.Millisecond {
		return baseWindow * 2
	}
	if rtt > 200*time.Millisecond {
		return baseWindow / 2
	}
	return baseWindow
}
