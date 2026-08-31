package tagger

import "testing"

func TestHasXingHeader(t *testing.T) {
	if !HasXingHeader([]byte("...Xing...")) {
		t.Errorf("expected true for Xing header")
	}
}
