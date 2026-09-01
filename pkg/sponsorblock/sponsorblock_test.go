package sponsorblock

import "testing"

func TestShouldSkipTime(t *testing.T) {
	segs := []SkipSegment{
		{Category: Sponsor, StartTime: 60.0, EndTime: 120.0},
	}
	if !ShouldSkipTime(segs, 90.0) {
		t.Errorf("expected skip at 90s")
	}
	if ShouldSkipTime(segs, 150.0) {
		t.Errorf("expected no skip at 150s")
	}
}
