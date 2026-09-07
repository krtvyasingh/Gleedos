package loudness

type BiquadCoefficients struct {
	B0, B1, B2, A1, A2 float64
}

func GetStage1Coefficients() BiquadCoefficients {
	return BiquadCoefficients{
		B0: 1.53512485958697,
		B1: -2.69169618940638,
		B2: 1.19839281085285,
		A1: -1.69065929318241,
		A2: 0.73248077421585,
	}
}
