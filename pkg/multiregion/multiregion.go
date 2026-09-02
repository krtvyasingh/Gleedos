package multiregion

import "hash/fnv"

type HashRing struct {
	regions []string
}

func NewHashRing(regions []string) *HashRing {
	return &HashRing{regions: regions}
}

func (h *HashRing) GetRegion(key string) string {
	if len(h.regions) == 0 {
		return ""
	}
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	idx := hasher.Sum32() % uint32(len(h.regions))
	return h.regions[idx]
}
