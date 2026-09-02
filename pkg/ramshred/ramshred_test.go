package ramshred

import "testing"

func TestRAMBufferShred(t *testing.T) {
	rb := NewRAMBuffer(1024)
	rb.Shred()
	for _, b := range rb.data {
		if b != 0 {
			t.Fatalf("buffer not zeroed after shred")
		}
	}
}
