package hls

import (
	"fmt"
	"strconv"
	"strings"
)

type ByteRange struct {
	Length int64
	Offset int64
}

func ParseByteRange(line string, lastOffset int64) (*ByteRange, error) {
	parts := strings.Split(strings.TrimPrefix(line, "#EXT-X-BYTERANGE:"), "@")
	length, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid byterange length: %w", err)
	}
	offset := lastOffset
	if len(parts) > 1 {
		off, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err == nil {
			offset = off
		}
	}
	return &ByteRange{Length: length, Offset: offset}, nil
}
