package iouring

type SQE struct {
	Opcode uint8
	Fd     int
	Buf    []byte
	Offset int64
}

type CQE struct {
	Res   int32
	Flags uint32
}

type RingSimulator struct {
	queue []SQE
}

func NewRingSimulator(capacity int) *RingSimulator {
	return &RingSimulator{queue: make([]SQE, 0, capacity)}
}

func (r *RingSimulator) Submit(sqe SQE) CQE {
	r.queue = append(r.queue, sqe)
	return CQE{Res: int32(len(sqe.Buf)), Flags: 0}
}
