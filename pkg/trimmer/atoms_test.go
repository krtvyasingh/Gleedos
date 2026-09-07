package trimmer

import "testing"

func TestFindAtomOffset(t *testing.T) {
	data := []byte("\x00\x00\x00\x20ftypmp42\x00\x00\x00\x10moov\x00\x00\x00\x08mdat")
	idx := FindAtomOffset(data, "moov")
	if idx == -1 {
		t.Errorf("failed to find moov atom")
	}
}
