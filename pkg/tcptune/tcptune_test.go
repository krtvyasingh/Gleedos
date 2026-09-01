package tcptune

import "testing"

func TestCalculateMicroChunkSize(t *testing.T) {
	chunk := CalculateMicroChunkSize(10*1024*1024, 50)
	if chunk < 256*1024 || chunk > 16*1024*1024 {
		t.Errorf("unexpected microchunk size: %d", chunk)
	}
}
