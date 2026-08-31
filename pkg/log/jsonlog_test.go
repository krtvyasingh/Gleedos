package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestJSONLogger(t *testing.T) {
	var buf bytes.Buffer
	l := NewJSONLogger(&buf)
	l.Log("INFO", "Download started")
	if !strings.Contains(buf.String(), `"level":"INFO"`) {
		t.Errorf("unexpected log output: %s", buf.String())
	}
}
