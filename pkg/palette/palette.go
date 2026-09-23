package palette

import "fmt"

type ColorHex string

func FormatRGBHex(r, g, b uint8) ColorHex {
	return ColorHex(fmt.Sprintf("#%02X%02X%02X", r, g, b))
}
