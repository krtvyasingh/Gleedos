package demux

import "bytes"

var EBMLHeader = []byte{0x1A, 0x45, 0xDF, 0xA3}

func IsMKV(header []byte) bool {
	return len(header) >= 4 && bytes.Equal(header[:4], EBMLHeader)
}
