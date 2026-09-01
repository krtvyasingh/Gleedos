package thumbnailer

import "testing"

func TestSelectBestFrame(t *testing.T) {
	frames := []FrameCandidate{
		{TimestampSec: 1.0, Contrast: 5.0, HasFace: false},
		{TimestampSec: 5.0, Contrast: 3.0, HasFace: true},
	}
	best := SelectBestFrame(frames)
	if best.TimestampSec != 5.0 {
		t.Errorf("expected frame with face to be selected, got %+v", best)
	}
}
