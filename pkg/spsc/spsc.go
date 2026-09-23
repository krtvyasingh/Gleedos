package spsc

import "sync/atomic"

type RingBuffer struct {
	buffer []byte
	mask   uint64
	head   uint64
	tail   uint64
}

func NewRingBuffer(sizePowerOfTwo int) *RingBuffer {
	size := 1 << sizePowerOfTwo
	return &RingBuffer{
		buffer: make([]byte, size),
		mask:   uint64(size - 1),
	}
}

func (r *RingBuffer) Push(b byte) bool {
	h := atomic.LoadUint64(&r.head)
	t := atomic.LoadUint64(&r.tail)
	if h-t > r.mask {
		return false // full
	}
	r.buffer[h&r.mask] = b
	atomic.StoreUint64(&r.head, h+1)
	return true
}

func (r *RingBuffer) Pop() (byte, bool) {
	h := atomic.LoadUint64(&r.head)
	t := atomic.LoadUint64(&r.tail)
	if h == t {
		return 0, false // empty
	}
	b := r.buffer[t&r.mask]
	atomic.StoreUint64(&r.tail, t+1)
	return b, true
}
