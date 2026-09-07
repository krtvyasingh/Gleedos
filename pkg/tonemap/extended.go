package tonemap

func ReinhardExtended(l, lWhite float64) float64 {
	if l <= 0 {
		return 0
	}
	if lWhite <= 0 {
		lWhite = 1.0
	}
	numerator := l * (1.0 + (l / (lWhite * lWhite)))
	return numerator / (1.0 + l)
}
