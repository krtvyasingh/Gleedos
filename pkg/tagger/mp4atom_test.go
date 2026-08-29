package tagger

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestParseAtoms(t *testing.T) {
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.BigEndian, uint32(16))
	buf.WriteString("ftypmp42")
	buf.Write([]byte{0, 0, 0, 0})

	atoms, err := ParseAtoms(&buf)
	if err != nil {
		t.Fatalf("ParseAtoms failed: %v", err)
	}

	if len(atoms) != 1 || atoms[0].Type != "ftyp" || atoms[0].Size != 16 {
		t.Errorf("unexpected atoms: %+v", atoms)
	}
}
