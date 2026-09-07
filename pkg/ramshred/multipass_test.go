package ramshred

import "testing"

func TestMultiPassShred(t *testing.T) {
	buf := []byte{1, 2, 3, 4, 5}
	MultiPassShred(buf)
	for _, b := range buf {
		if b != 0x00 {
			t.Errorf("memory not zeroed")
		}
	}
}
