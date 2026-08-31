package nativehost

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

type NativeMessage struct {
	URL string `json:"url"`
}

func ReadNativeMessage(r io.Reader) (*NativeMessage, error) {
	var length uint32
	if err := binary.Read(r, binary.NativeEndian, &length); err != nil {
		return nil, err
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	var msg NativeMessage
	if err := json.Unmarshal(buf, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
