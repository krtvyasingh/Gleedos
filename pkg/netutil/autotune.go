package netutil

func CalculateOptimalThreads(fileSizeBytes int64) int {
	if fileSizeBytes < 10*1024*1024 {
		return 2
	}
	if fileSizeBytes < 100*1024*1024 {
		return 4
	}
	if fileSizeBytes < 1024*1024*1024 {
		return 8
	}
	return 16
}
