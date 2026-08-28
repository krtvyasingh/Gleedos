package webhook

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDispatch(t *testing.T) {
	received := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = true
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ev := Event{
		Event:     "download.completed",
		URL:       "https://example.com/video",
		Path:      "/tmp/video.mp4",
		Success:   true,
		Timestamp: time.Now(),
	}

	if err := Dispatch(context.Background(), server.URL, ev); err != nil {
		t.Fatalf("Dispatch failed: %v", err)
	}
	if !received {
		t.Errorf("webhook was not received")
	}
}
