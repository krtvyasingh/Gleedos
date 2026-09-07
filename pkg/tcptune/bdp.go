package tcptune

func CalculateOptimalWindow(bandwidthBps int64, rttMs int64) int {
	if bandwidthBps <= 0 || rttMs <= 0 {
		return 64 * 1024 // 64KB default
	}
	bdp := (bandwidthBps * rttMs) / 8000
	if bdp < 64*1024 {
		return 64 * 1024
	}
	if bdp > 32*1024*1024 {
		return 32 * 1024 * 1024 // 32MB cap
	}
	return int(bdp)
}
