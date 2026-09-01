package remuxer

import (
	"bytes"
	"testing"
)

func TestRemuxMP4(t *testing.T) {
	var out bytes.Buffer
	err := RemuxMP4([]byte("video_data"), []byte("audio_data"), &out)
	if err != nil || out.Len() == 0 {
		t.Fatalf("RemuxMP4 failed: %v", err)
	}
}
