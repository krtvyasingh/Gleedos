package thumbnailer

type FrameCandidate struct {
	TimestampSec float64
	Contrast     float64
	HasFace      bool
}

func SelectBestFrame(frames []FrameCandidate) FrameCandidate {
	if len(frames) == 0 {
		return FrameCandidate{}
	}
	best := frames[0]
	var bestScore float64
	for _, f := range frames {
		score := f.Contrast
		if f.HasFace {
			score += 10.0
		}
		if score > bestScore {
			bestScore = score
			best = f
		}
	}
	return best
}
