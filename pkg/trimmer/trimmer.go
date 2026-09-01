package trimmer

import (
	"io"
	"time"
)

type KeyframeInfo struct {
	Timestamp time.Duration
	ByteOffset int64
}

func FindNearestKeyframe(keyframes []KeyframeInfo, target time.Duration) KeyframeInfo {
	if len(keyframes) == 0 {
		return KeyframeInfo{}
	}
	best := keyframes[0]
	minDiff := absDuration(keyframes[0].Timestamp - target)
	for _, k := range keyframes {
		diff := absDuration(k.Timestamp - target)
		if diff < minDiff {
			minDiff = diff
			best = k
		}
	}
	return best
}

func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}

func SliceStreamLossless(r io.ReaderAt, startOff, length int64, w io.Writer) (int64, error) {
	sr := io.NewSectionReader(r, startOff, length)
	return io.Copy(w, sr)
}
