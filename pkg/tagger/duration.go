package tagger

import "time"

func EstimateDuration(sizeBytes int64, bitrateKbps int) time.Duration {
	if bitrateKbps <= 0 {
		return 0
	}
	bytesPerSec := int64(bitrateKbps * 1000 / 8)
	return time.Duration(sizeBytes/bytesPerSec) * time.Second
}
