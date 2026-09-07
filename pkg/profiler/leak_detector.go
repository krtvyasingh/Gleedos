package profiler

import "runtime"

type LeakDetector struct {
	baseline int
}

func NewLeakDetector() *LeakDetector {
	return &LeakDetector{baseline: runtime.NumGoroutine()}
}

func (l *LeakDetector) DetectDrift(tolerance int) (bool, int) {
	curr := runtime.NumGoroutine()
	drift := curr - l.baseline
	return drift > tolerance, drift
}
