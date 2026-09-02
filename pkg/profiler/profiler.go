package profiler

import (
	"fmt"
	"runtime"
)

type RuntimeStats struct {
	AllocMB     uint64
	NumGoroutine int
}

func CaptureStats() RuntimeStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return RuntimeStats{
		AllocMB:      m.Alloc / (1024 * 1024),
		NumGoroutine: runtime.NumGoroutine(),
	}
}

func FormatStats(s RuntimeStats) string {
	return fmt.Sprintf("Alloc: %d MB | Goroutines: %d", s.AllocMB, s.NumGoroutine)
}
