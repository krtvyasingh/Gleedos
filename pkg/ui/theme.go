package ui

const (
	ColorReset  = "[0m"
	ColorGreen  = "[32m"
	ColorYellow = "[33m"
	ColorBlue   = "[34m"
	ColorCyan   = "[36m"
)

func Colorize(color, text string) string {
	return color + text + ColorReset
}
