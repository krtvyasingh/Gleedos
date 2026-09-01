package parity

import (
	"crypto/sha256"
	"encoding/hex"
)

type ParityBlock struct {
	Index int
	Hash  string
}

func GenerateParity(data []byte, blockSize int) []ParityBlock {
	if blockSize <= 0 {
		blockSize = 64 * 1024
	}
	var blocks []ParityBlock
	idx := 0
	for i := 0; i < len(data); i += blockSize {
		end := i + blockSize
		if end > len(data) {
			end = len(data)
		}
		h := sha256.Sum256(data[i:end])
		blocks = append(blocks, ParityBlock{
			Index: idx,
			Hash:  hex.EncodeToString(h[:]),
		})
		idx++
	}
	return blocks
}
