package bitset

type ChunkBitset struct {
	bits []uint64
	size int
}

func NewChunkBitset(totalChunks int) *ChunkBitset {
	numWords := (totalChunks + 63) / 64
	return &ChunkBitset{bits: make([]uint64, numWords), size: totalChunks}
}

func (c *ChunkBitset) Set(idx int) {
	if idx >= 0 && idx < c.size {
		c.bits[idx/64] |= (1 << (idx % 64))
	}
}

func (c *ChunkBitset) IsSet(idx int) bool {
	if idx >= 0 && idx < c.size {
		return (c.bits[idx/64] & (1 << (idx % 64))) != 0
	}
	return false
}
