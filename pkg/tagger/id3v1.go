package tagger

import "bytes"

func HasID3v1(data []byte) bool {
	return len(data) >= 128 && bytes.Equal(data[len(data)-128:len(data)-125], []byte("TAG"))
}
