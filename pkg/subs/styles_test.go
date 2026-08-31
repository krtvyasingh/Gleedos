package subs

import "testing"

func TestCreateASSStyle(t *testing.T) {
	st := CreateASSStyle("Custom", "Helvetica", 24)
	if st != "Style: Custom,Helvetica" {
		t.Errorf("unexpected style output: %s", st)
	}
}
