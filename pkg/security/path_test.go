package security

import "testing"

func TestIsSafePath(t *testing.T) {
	_, err := IsSafePath("/app/downloads", "../../etc/passwd")
	if err == nil {
		t.Errorf("expected error on path traversal")
	}
	safe, err := IsSafePath("/app/downloads", "videos/test.mp4")
	if err != nil || safe != "/app/downloads/videos/test.mp4" {
		t.Errorf("unexpected safe path: %s, %v", safe, err)
	}
}
