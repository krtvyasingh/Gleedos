package log

import (
	"encoding/json"
	"io"
	"time"
)

type LogEntry struct {
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type JSONLogger struct {
	out io.Writer
}

func NewJSONLogger(w io.Writer) *JSONLogger {
	return &JSONLogger{out: w}
}

func (l *JSONLogger) Log(level, msg string) {
	entry := LogEntry{Level: level, Message: msg, Timestamp: time.Now()}
	_ = json.NewEncoder(l.out).Encode(entry)
}
