package log

import "testing"

func TestNewSpan(t *testing.T) {
	s := NewSpan("http.get", "trace_1", "span_1")
	if s.Name != "http.get" || s.TraceID != "trace_1" {
		t.Errorf("unexpected span: %+v", s)
	}
}
