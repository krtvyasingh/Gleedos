package audioid

type ConstellationPoint struct {
	FreqBin  int
	TimeSlot int
}

func ExtractPeaks(spectrum [][]float64, threshold float64) []ConstellationPoint {
	var points []ConstellationPoint
	for t, row := range spectrum {
		for f, mag := range row {
			if mag > threshold {
				points = append(points, ConstellationPoint{FreqBin: f, TimeSlot: t})
			}
		}
	}
	return points
}
