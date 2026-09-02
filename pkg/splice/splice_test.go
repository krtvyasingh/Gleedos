package splice

import (
	"bytes"
	"strings"
	"testing"
)

func TestZeroCopySplice(t *testing.T) {
	src := strings.NewReader("splice_payload")
	var dst bytes.Buffer
	n, err := ZeroCopySplice(src, &dst)
	if err != nil || n != 14 {
		t.Fatalf("ZeroCopySplice failed: %v, %d", err, n)
	}
}
