package mp4zero

import (
	"bytes"
	"testing"
)

func TestWriteBoxHeader(t *testing.T) {
	var buf bytes.Buffer
	err := WriteBoxHeader(&buf, [4]byte{'f', 't', 'y', 'p'}, 16)
	if err != nil || buf.Len() != 8 {
		t.Fatalf("WriteBoxHeader failed: %v, len: %d", err, buf.Len())
	}
}
