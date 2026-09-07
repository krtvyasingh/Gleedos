package demux

type FLACBlockHeader struct {
	IsLast    bool
	BlockType int
	Length    int
}

func ParseFLACBlockHeader(buf []byte) FLACBlockHeader {
	if len(buf) < 4 {
		return FLACBlockHeader{}
	}
	isLast := (buf[0] & 0x80) != 0
	bType := int(buf[0] & 0x7F)
	lenVal := int(buf[1])<<16 | int(buf[2])<<8 | int(buf[3])
	return FLACBlockHeader{IsLast: isLast, BlockType: bType, Length: lenVal}
}
