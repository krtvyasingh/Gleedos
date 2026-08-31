package storage

type RingBuffer struct {
	buf  []byte
	size int
	head int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{buf: make([]byte, size), size: size}
}

func (r *RingBuffer) Write(p []byte) int {
	for i, b := range p {
		r.buf[(r.head+i)%r.size] = b
	}
	r.head = (r.head + len(p)) % r.size
	return len(p)
}
