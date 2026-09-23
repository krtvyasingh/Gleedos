package simdhash

import "hash/crc32"

var table = crc32.MakeTable(crc32.Castagnoli)

func ChecksumFast(data []byte) uint32 {
	return crc32.Checksum(data, table)
}
