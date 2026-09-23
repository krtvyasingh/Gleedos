package shm

type SHMBuffer struct {
	SegmentID string
	Size      int
}

func NewSHMBuffer(segmentID string, size int) *SHMBuffer {
	return &SHMBuffer{SegmentID: segmentID, Size: size}
}
