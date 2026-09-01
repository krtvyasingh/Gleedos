package pipe

import (
	"bytes"
	"io"
	"testing"
)

func TestPipeStream(t *testing.T) {
	src := bytes.NewReader([]byte("stream_chunk"))
	var out bytes.Buffer
	n, err := io.Copy(&out, src)
	if err != nil || n != int64(len("stream_chunk")) {
		t.Fatalf("io.Copy failed: %v", err)
	}
}
