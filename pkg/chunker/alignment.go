package chunker

func AlignChunkBoundary(offset int64, blockSize int64) int64 {
	if blockSize <= 0 {
		return offset
	}
	remainder := offset % blockSize
	if remainder == 0 {
		return offset
	}
	return offset - remainder
}
