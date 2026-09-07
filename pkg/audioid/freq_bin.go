package audioid

func NormalizeBin(value, minVal, maxVal float64) float64 {
	if maxVal <= minVal {
		return 0.0
	}
	norm := (value - minVal) / (maxVal - minVal)
	if norm < 0.0 {
		return 0.0
	}
	if norm > 1.0 {
		return 1.0
	}
	return norm
}
