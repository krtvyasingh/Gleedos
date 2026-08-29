package tagger

import (
	"encoding/binary"
	"io"
	"os"
)

type Atom struct {
	Type string
	Size uint32
}

func ParseAtoms(r io.Reader) ([]Atom, error) {
	var atoms []Atom
	buf := make([]byte, 8)

	for {
		_, err := io.ReadFull(r, buf)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			return atoms, err
		}

		size := binary.BigEndian.Uint32(buf[0:4])
		atomType := string(buf[4:8])
		atoms = append(atoms, Atom{Type: atomType, Size: size})

		if size > 8 {
			if _, err := io.CopyN(io.Discard, r, int64(size-8)); err != nil {
				break
			}
		}
	}
	return atoms, nil
}

func HasMoovAtom(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	atoms, err := ParseAtoms(f)
	if err != nil {
		return false
	}

	for _, a := range atoms {
		if a.Type == "moov" || a.Type == "ftyp" {
			return true
		}
	}
	return false
}
