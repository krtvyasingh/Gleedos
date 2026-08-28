package pool

import "sync"

const DefaultBufferSize = 32 * 1024 // 32KB buffer

var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, DefaultBufferSize)
		return &b
	},
}

func GetBuffer() *[]byte {
	return bufPool.Get().(*[]byte)
}

func PutBuffer(b *[]byte) {
	if b != nil && len(*b) == DefaultBufferSize {
		bufPool.Put(b)
	}
}
