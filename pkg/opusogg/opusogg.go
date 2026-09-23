package opusogg

type OpusHead struct {
	Channels   uint8
	PreSkip    uint16
	SampleRate uint32
}

func CreateOpusHead(channels uint8, sampleRate uint32) []byte {
	head := make([]byte, 19)
	copy(head[0:8], "OpusHead")
	head[8] = 1 // version
	head[9] = channels
	head[10] = 0 // preskip low
	head[11] = 0 // preskip high
	head[12] = byte(sampleRate)
	head[13] = byte(sampleRate >> 8)
	head[14] = byte(sampleRate >> 16)
	head[15] = byte(sampleRate >> 24)
	return head
}
