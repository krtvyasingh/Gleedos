package actorindex

import (
	"testing"
	"time"
)

func TestMergeIntervals(t *testing.T) {
	apps := []ActorAppearance{
		{ActorName: "Alice", Start: 0, End: 5 * time.Second},
		{ActorName: "Alice", Start: 5 * time.Second, End: 10 * time.Second},
	}
	merged := MergeIntervals(apps)
	if len(merged) != 1 || merged[0].End != 10*time.Second {
		t.Errorf("merge failed: %+v", merged)
	}
}
