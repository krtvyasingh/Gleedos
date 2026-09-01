package remuxer

import (
	"bytes"
	"encoding/binary"
	"io"
)

type TrackData struct {
	Type string
	Payload []byte
}

func RemuxMP4(videoTrack, audioTrack []byte, w io.Writer) error {
	// Write ftyp box
	ftyp := []byte("ftypmp42\x00\x00\x00\x00isommp42")
	_ = binary.Write(w, binary.BigEndian, uint32(len(ftyp)+4))
	_, _ = w.Write(ftyp)

	// Write mdat payload
	var mdat bytes.Buffer
	mdat.WriteString("mdat")
	mdat.Write(videoTrack)
	mdat.Write(audioTrack)
	_ = binary.Write(w, binary.BigEndian, uint32(mdat.Len()+4))
	_, err := io.Copy(w, &mdat)
	return err
}
