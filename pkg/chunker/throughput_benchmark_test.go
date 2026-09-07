package chunker

import (
	"bytes"
	"io"
	"testing"
)

func BenchmarkBufferThroughput(b *testing.B) {
	data := make([]byte, 64*1024)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		_, _ = io.Copy(io.Discard, r)
	}
}
