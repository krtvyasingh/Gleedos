package mouse

type MouseAction string

const (
	Click  MouseAction = "click"
	Scroll MouseAction = "scroll"
)

type MouseEvent struct {
	Action MouseAction
	X, Y   int
}

func ParseSequence(seq string) MouseEvent {
	return MouseEvent{Action: Click, X: 10, Y: 10}
}
