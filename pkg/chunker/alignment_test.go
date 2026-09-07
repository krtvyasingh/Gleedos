package chunker

import "testing"

func TestAlignChunkBoundary(t *testing.T) {
	aligned := AlignChunkBoundary(1050, 512)
	if aligned != 1024 {
		t.Errorf("expected 1024, got %d", aligned)
	}
}
