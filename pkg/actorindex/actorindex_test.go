package actorindex

import (
	"testing"
	"time"
)

func TestGroupAppearances(t *testing.T) {
	apps := []ActorAppearance{
		{ActorName: "Host", Start: 0, End: 10 * time.Second},
		{ActorName: "Host", Start: 20 * time.Second, End: 30 * time.Second},
	}
	counts := GroupAppearances(apps)
	if counts["Host"] != 2 {
		t.Errorf("expected 2 appearances, got %d", counts["Host"])
	}
}
