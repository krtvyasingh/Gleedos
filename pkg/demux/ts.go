package demux

import (
	"errors"
	"io"
)

const SyncByte = 0x47
const PacketSize = 188

type TSPacket struct {
	PID     uint16
	Payload []byte
}

func ReadTSPacket(r io.Reader) (*TSPacket, error) {
	buf := make([]byte, PacketSize)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	if buf[0] != SyncByte {
		return nil, errors.New("invalid TS sync byte")
	}

	pid := uint16(buf[1]&0x1F)<<8 | uint16(buf[2])
	return &TSPacket{
		PID:     pid,
		Payload: buf[4:],
	}, nil
}
