package subs

import "testing"

func TestIsMicroDVD(t *testing.T) {
	if !IsMicroDVD("{100}{200}Hello") {
		t.Errorf("expected true for MicroDVD line")
	}
}
