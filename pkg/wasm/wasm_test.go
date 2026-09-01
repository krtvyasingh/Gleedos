package wasm

import "testing"

func TestCalculateChunkBoundaries(t *testing.T) {
	ranges := CalculateChunkBoundaries(100, 4)
	if len(ranges) != 4 || ranges[3][1] != 99 {
		t.Errorf("unexpected chunk ranges: %+v", ranges)
	}
}
