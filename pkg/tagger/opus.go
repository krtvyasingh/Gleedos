package tagger

import "bytes"

func IsOpusHead(data []byte) bool {
	return bytes.HasPrefix(data, []byte("OpusHead"))
}
