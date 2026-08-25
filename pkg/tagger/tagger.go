package tagger

import (
	"bytes"
	"errors"
	"io"
	"os"
)

// MediaType identifies common media file formats.
type MediaType string

const (
	TypeMP4     MediaType = "mp4"
	TypeMKV     MediaType = "mkv"
	TypeWebM    MediaType = "webm"
	TypeMP3     MediaType = "mp3"
	TypeTS      MediaType = "ts"
	TypeUnknown MediaType = "unknown"
)

// DetectMediaType analyzes initial file bytes to identify the container format.
func DetectMediaType(r io.Reader) (MediaType, error) {
	header := make([]byte, 512)
	n, err := io.ReadFull(r, header)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) && !errors.Is(err, io.EOF) {
		return TypeUnknown, err
	}
	if n < 4 {
		return TypeUnknown, errors.New("file too small")
	}

	// MP4 / M4A: check for 'ftyp' at offset 4
	if n >= 8 && string(header[4:8]) == "ftyp" {
		return TypeMP4, nil
	}

	// MKV / WebM: EBML header 0x1A 0x45 0xDF 0xA3
	if bytes.HasPrefix(header, []byte{0x1A, 0x45, 0xDF, 0xA3}) {
		if bytes.Contains(header[:n], []byte("webm")) {
			return TypeWebM, nil
		}
		return TypeMKV, nil
	}

	// MP3: ID3 header or sync word 0xFF 0xFB / 0xFF 0xF3
	if bytes.HasPrefix(header, []byte("ID3")) || (header[0] == 0xFF && (header[1]&0xE0) == 0xE0) {
		return TypeMP3, nil
	}

	// MPEG-TS: Sync byte 0x47
	if header[0] == 0x47 {
		return TypeTS, nil
	}

	return TypeUnknown, nil
}

// ValidateMediaFile ensures the file exists, is non-empty, and has a recognizable media signature.
func ValidateMediaFile(path string) (MediaType, error) {
	f, err := os.Open(path)
	if err != nil {
		return TypeUnknown, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return TypeUnknown, err
	}
	if stat.Size() == 0 {
		return TypeUnknown, errors.New("media file is 0 bytes")
	}

	return DetectMediaType(f)
}

// Metadata holds track metadata.
type Metadata struct {
	Title   string
	Artist  string
	Comment string
}

// InjectID3Tag prepends an ID3v2.3 tag with Title, Artist, and Comment to an MP3 file.
func InjectID3Tag(filePath string, meta Metadata) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Build ID3v2.3 tag buffer
	var frames bytes.Buffer

	addFrame := func(id string, text string) {
		if text == "" {
			return
		}
		frameBody := append([]byte{0x00}, []byte(text)...) // ISO-8859-1 encoding flag + text
		frameSize := len(frameBody)

		frames.WriteString(id)
		// 4 bytes frame size
		frames.WriteByte(byte(frameSize >> 24))
		frames.WriteByte(byte(frameSize >> 16))
		frames.WriteByte(byte(frameSize >> 8))
		frames.WriteByte(byte(frameSize))
		// 2 flag bytes
		frames.Write([]byte{0x00, 0x00})
		frames.Write(frameBody)
	}

	addFrame("TIT2", meta.Title)
	addFrame("TPE1", meta.Artist)
	if meta.Comment != "" {
		// COMM frame: 1 byte encoding, 3 bytes lang ('eng'), 1 byte short desc (0x00), comment text
		commBody := append([]byte{0x00, 'e', 'n', 'g', 0x00}, []byte(meta.Comment)...)
		frames.WriteString("COMM")
		sz := len(commBody)
		frames.WriteByte(byte(sz >> 24))
		frames.WriteByte(byte(sz >> 16))
		frames.WriteByte(byte(sz >> 8))
		frames.WriteByte(byte(sz))
		frames.Write([]byte{0x00, 0x00})
		frames.Write(commBody)
	}

	tagPayload := frames.Bytes()
	tagLen := len(tagPayload)

	// ID3v2 Header: ID3 + 0x03 0x00 (v2.3) + flags (0x00) + 4-byte synchsafe integer size
	var id3Header bytes.Buffer
	id3Header.WriteString("ID3")
	id3Header.Write([]byte{0x03, 0x00, 0x00})

	// Encode synchsafe integer
	id3Header.WriteByte(byte((tagLen >> 21) & 0x7F))
	id3Header.WriteByte(byte((tagLen >> 14) & 0x7F))
	id3Header.WriteByte(byte((tagLen >> 7) & 0x7F))
	id3Header.WriteByte(byte(tagLen & 0x7F))
	id3Header.Write(tagPayload)

	// Prepend tag to existing file data
	newData := append(id3Header.Bytes(), data...)
	return os.WriteFile(filePath, newData, 0644)
}
