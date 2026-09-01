package pipe

import (
	"io"
	"os"
)

func StreamToStdout(r io.Reader) (int64, error) {
	return io.Copy(os.Stdout, r)
}
