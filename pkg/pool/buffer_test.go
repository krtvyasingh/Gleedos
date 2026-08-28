package pool

import "testing"

func TestBufferPool(t *testing.T) {
	buf := GetBuffer()
	if buf == nil || len(*buf) != DefaultBufferSize {
		t.Fatalf("unexpected buffer size")
	}
	PutBuffer(buf)
}

func BenchmarkBufferPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := GetBuffer()
			PutBuffer(buf)
		}
	})
}
