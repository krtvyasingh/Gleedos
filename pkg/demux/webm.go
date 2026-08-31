package demux

import "bytes"

func IsWebM(header []byte) bool {
	return IsMKV(header) && bytes.Contains(header, []byte("webm"))
}
