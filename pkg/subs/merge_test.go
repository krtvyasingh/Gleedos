package subs

import "testing"

func TestCombineTracks(t *testing.T) {
	tracks := []SubTrack{{Language: "en"}, {Language: "es"}}
	if CombineTracks(tracks) != 2 {
		t.Errorf("expected 2 tracks")
	}
}
