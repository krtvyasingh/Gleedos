package trimmer

import "bytes"

func FindAtomOffset(data []byte, atomType string) int64 {
	if len(atomType) != 4 {
		return -1
	}
	idx := bytes.Index(data, []byte(atomType))
	return int64(idx)
}
