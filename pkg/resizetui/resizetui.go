package resizetui

type Dimensions struct {
	Width  int
	Height int
}

func ClampDimensions(w, h int) Dimensions {
	if w < 40 {
		w = 40
	}
	if h < 10 {
		h = 10
	}
	return Dimensions{Width: w, Height: h}
}
