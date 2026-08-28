package chunker

import "os"

// PWriteChunk executes offset write using direct OS write.
func PWriteChunk(f *os.File, b []byte, off int64) (int, error) {
	return f.WriteAt(b, off)
}
