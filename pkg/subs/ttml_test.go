package subs

import "testing"

func TestIsTTML(t *testing.T) {
	if !IsTTML("<tt xmlns=\"http://www.w3.org/ns/ttml\">") {
		t.Errorf("expected true for TTML")
	}
}
