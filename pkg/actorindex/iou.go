package actorindex

type BoundingBox struct {
	X, Y, W, H float64
}

func CalculateIoU(b1, b2 BoundingBox) float64 {
	x1 := max(b1.X, b2.X)
	y1 := max(b1.Y, b2.Y)
	x2 := min(b1.X+b1.W, b2.X+b2.W)
	y2 := min(b1.Y+b1.H, b2.Y+b2.H)

	interArea := max(0, x2-x1) * max(0, y2-y1)
	b1Area := b1.W * b1.H
	b2Area := b2.W * b2.H
	unionArea := b1Area + b2Area - interArea
	if unionArea <= 0 {
		return 0
	}
	return interArea / unionArea
}
