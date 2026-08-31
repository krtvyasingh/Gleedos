package ui

var SpinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func GetSpinnerFrame(tick int) string {
	return SpinnerFrames[tick%len(SpinnerFrames)]
}
