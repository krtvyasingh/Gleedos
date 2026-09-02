package ramshred

import (
	"crypto/rand"
	"io"
)

type RAMBuffer struct {
	data []byte
}

func NewRAMBuffer(size int) *RAMBuffer {
	return &RAMBuffer{data: make([]byte, size)}
}

func (r *RAMBuffer) Shred() {
	_, _ = io.ReadFull(rand.Reader, r.data)
	for i := range r.data {
		r.data[i] = 0
	}
}
