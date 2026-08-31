package nativehost

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestReadNativeMessage(t *testing.T) {
	payload := []byte(`{"url": "https://example.com/video"}`)
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.NativeEndian, uint32(len(payload)))
	buf.Write(payload)

	msg, err := ReadNativeMessage(&buf)
	if err != nil || msg.URL != "https://example.com/video" {
		t.Fatalf("ReadNativeMessage failed: %v, %+v", err, msg)
	}
}
