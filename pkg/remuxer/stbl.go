package remuxer

import (
	"bytes"
	"encoding/binary"
)

func BuildSTBLBox(sampleCount uint32) []byte {
	var buf bytes.Buffer
	buf.WriteString("stbl")
	_ = binary.Write(&buf, binary.BigEndian, sampleCount)
	return buf.Bytes()
}
