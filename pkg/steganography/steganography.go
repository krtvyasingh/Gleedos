package steganography

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var ChaffMarker = []byte("GCHAF")

func HidePayload(mp4Data, secret []byte) []byte {
	var buf bytes.Buffer
	buf.Write(mp4Data)
	buf.Write(ChaffMarker)
	_ = binary.Write(&buf, binary.BigEndian, uint32(len(secret)))
	buf.Write(secret)
	return buf.Bytes()
}

func ExtractPayload(mp4Data []byte) ([]byte, error) {
	idx := bytes.LastIndex(mp4Data, ChaffMarker)
	if idx == -1 {
		return nil, errors.New("no hidden payload found")
	}
	rem := mp4Data[idx+len(ChaffMarker):]
	if len(rem) < 4 {
		return nil, errors.New("corrupted chaff header")
	}
	length := binary.BigEndian.Uint32(rem[:4])
	return rem[4 : 4+length], nil
}
