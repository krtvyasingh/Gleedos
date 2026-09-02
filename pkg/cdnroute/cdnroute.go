package cdnroute

type EdgeMirror struct {
	IP        string
	LatencyMs float64
	LossRate  float64
}

func ScoreMirror(m EdgeMirror) float64 {
	return m.LatencyMs * (1.0 + m.LossRate*10.0)
}

func SelectOptimalEdge(mirrors []EdgeMirror) EdgeMirror {
	if len(mirrors) == 0 {
		return EdgeMirror{}
	}
	best := mirrors[0]
	bestScore := ScoreMirror(mirrors[0])
	for _, m := range mirrors {
		score := ScoreMirror(m)
		if score < bestScore {
			bestScore = score
			best = m
		}
	}
	return best
}
