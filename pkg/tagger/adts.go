package tagger

func IsADTS(hdr []byte) bool {
	if len(hdr) < 2 {
		return false
	}
	return hdr[0] == 0xFF && (hdr[1]&0xF0) == 0xF0
}
