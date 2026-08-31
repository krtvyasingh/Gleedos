package subs

import "testing"

func TestIsSubViewerTimestamp(t *testing.T) {
	if !IsSubViewerTimestamp("0:00:01.000,0:00:04.000") {
		t.Errorf("expected true for SBV timestamp")
	}
}
