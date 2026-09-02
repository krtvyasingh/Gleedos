package tonemap

func ReinhardTonemap(hdrLinear float64) float64 {
	if hdrLinear <= 0 {
		return 0
	}
	return hdrLinear / (1.0 + hdrLinear)
}
