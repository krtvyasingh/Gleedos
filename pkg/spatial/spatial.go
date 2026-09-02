package spatial

type SurroundChannels struct {
	FL, FR, FC, LFE, SL, SR float64
}

func DownmixToBinaural(c SurroundChannels) (left, right float64) {
	left = c.FL + (0.707 * c.FC) + (0.5 * c.LFE) + (0.8 * c.SL)
	right = c.FR + (0.707 * c.FC) + (0.5 * c.LFE) + (0.8 * c.SR)
	return left, right
}
