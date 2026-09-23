package palette

import "testing"

func TestFormatRGBHex(t *testing.T) {
	hex := FormatRGBHex(255, 128, 0)
	if hex != "#FF8000" {
		t.Errorf("unexpected hex: %s", hex)
	}
}
