package security

import (
	"path/filepath"
	"testing"
)

func TestIsSafePath(t *testing.T) {
	_, err := IsSafePath("/app/downloads", "../../etc/passwd")
	if err == nil {
		t.Errorf("expected error on path traversal")
	}
	expected := filepath.Join(filepath.Clean("/app/downloads"), "videos", "test.mp4")
	safe, err := IsSafePath("/app/downloads", filepath.Join("videos", "test.mp4"))
	if err != nil || safe != expected {
		t.Errorf("unexpected safe path: %s (want %s), %v", safe, expected, err)
	}
}
