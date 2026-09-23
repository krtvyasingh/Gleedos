package opusogg

import (
	"bytes"
	"testing"
)

func TestCreateOpusHead(t *testing.T) {
	head := CreateOpusHead(2, 48000)
	if !bytes.HasPrefix(head, []byte("OpusHead")) || len(head) != 19 {
		t.Errorf("invalid OpusHead: %v", head)
	}
}
