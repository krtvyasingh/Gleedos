package storage

import (
	"crypto/sha256"
	"encoding/hex"
)

type DedupStore struct {
	seen map[string]bool
}

func NewDedupStore() *DedupStore {
	return &DedupStore{seen: make(map[string]bool)}
}

func (d *DedupStore) IsDuplicate(chunk []byte) bool {
	h := sha256.Sum256(chunk)
	key := hex.EncodeToString(h[:])
	if d.seen[key] {
		return true
	}
	d.seen[key] = true
	return false
}
