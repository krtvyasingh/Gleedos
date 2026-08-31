package log

type Span struct {
	Name    string
	TraceID string
	SpanID  string
}

func NewSpan(name, traceID, spanID string) Span {
	return Span{Name: name, TraceID: traceID, SpanID: spanID}
}
