package mp4zero

import (
	"encoding/binary"
	"io"
)

func WriteBoxHeader(w io.Writer, boxType [4]byte, payloadSize uint32) error {
	totalSize := payloadSize + 8
	var header [8]byte
	binary.BigEndian.PutUint32(header[0:4], totalSize)
	copy(header[4:8], boxType[:])
	_, err := w.Write(header[:])
	return err
}
