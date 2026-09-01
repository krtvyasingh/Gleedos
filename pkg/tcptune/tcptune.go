package tcptune

func CalculateMicroChunkSize(currentSpeedBps int64, latencyMs int64) int64 {
	if currentSpeedBps <= 0 {
		return 1024 * 1024 // 1MB fallback
	}
	bandwidthDelayProduct := (currentSpeedBps * latencyMs) / 1000
	if bandwidthDelayProduct < 256*1024 {
		return 256 * 1024
	}
	if bandwidthDelayProduct > 16*1024*1024 {
		return 16 * 1024 * 1024
	}
	return bandwidthDelayProduct
}
