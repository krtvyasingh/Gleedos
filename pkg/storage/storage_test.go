package storage

import (
	"context"
	"strings"
	"testing"
)

func TestLocalStorage(t *testing.T) {
	var s StorageProvider = &LocalStorage{}
	if s.Name() != "local" {
		t.Errorf("expected local provider name")
	}
	if err := s.Upload(context.Background(), "test.mp4", strings.NewReader("data")); err != nil {
		t.Errorf("upload failed: %v", err)
	}
}
