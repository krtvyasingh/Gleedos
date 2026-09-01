package trimmer

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestFindNearestKeyframe(t *testing.T) {
	kf := []KeyframeInfo{
		{Timestamp: 0 * time.Second, ByteOffset: 0},
		{Timestamp: 5 * time.Second, ByteOffset: 5000},
		{Timestamp: 10 * time.Second, ByteOffset: 10000},
	}
	best := FindNearestKeyframe(kf, 6*time.Second)
	if best.Timestamp != 5*time.Second {
		t.Errorf("expected 5s keyframe, got %v", best.Timestamp)
	}

	reader := strings.NewReader("0123456789ABCDEF")
	var buf bytes.Buffer
	n, err := SliceStreamLossless(reader, 4, 4, &buf)
	if err != nil || n != 4 || buf.String() != "4567" {
		t.Errorf("unexpected slice: %s (%d bytes, err: %v)", buf.String(), n, err)
	}
}
