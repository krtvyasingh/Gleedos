package tagger

func ParseADTSProfile(header []byte) int {
	if len(header) < 3 {
		return 0
	}
	if header[0] != 0xFF || (header[1]&0xF0) != 0xF0 {
		return 0
	}
	return int((header[2] & 0xC0) >> 6)
}
