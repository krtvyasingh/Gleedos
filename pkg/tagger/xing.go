package tagger

import "bytes"

func HasXingHeader(data []byte) bool {
	return bytes.Contains(data, []byte("Xing")) || bytes.Contains(data, []byte("Info"))
}
