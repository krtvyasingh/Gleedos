package remuxer

import (
	"bytes"
	"testing"
)

func TestBuildSTBLBox(t *testing.T) {
	box := BuildSTBLBox(100)
	if !bytes.HasPrefix(box, []byte("stbl")) {
		t.Errorf("invalid stbl box prefix")
	}
}
