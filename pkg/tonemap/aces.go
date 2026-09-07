package tonemap

func ACESFilmicTonemap(x float64) float64 {
	if x <= 0 {
		return 0
	}
	a := 2.51
	b := 0.03
	c := 2.43
	d := 0.59
	e := 0.14
	return (x*(a*x+b)) / (x*(c*x+d)+e)
}
