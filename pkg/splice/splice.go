package splice

import "io"

func ZeroCopySplice(src io.Reader, dst io.Writer) (int64, error) {
	return io.Copy(dst, src)
}
