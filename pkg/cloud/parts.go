package cloud

func CalculatePartCount(fileSizeBytes int64, partSizeBytes int64) int {
	if partSizeBytes <= 0 || fileSizeBytes <= 0 {
		return 1
	}
	parts := fileSizeBytes / partSizeBytes
	if fileSizeBytes%partSizeBytes != 0 {
		parts++
	}
	return int(parts)
}
