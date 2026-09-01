package wasm

func CalculateChunkBoundaries(totalSize int64, chunks int) [][2]int64 {
	if chunks <= 0 || totalSize <= 0 {
		return nil
	}
	chunkSize := totalSize / int64(chunks)
	var ranges [][2]int64
	for i := 0; i < chunks; i++ {
		start := int64(i) * chunkSize
		end := start + chunkSize - 1
		if i == chunks-1 {
			end = totalSize - 1
		}
		ranges = append(ranges, [2]int64{start, end})
	}
	return ranges
}
