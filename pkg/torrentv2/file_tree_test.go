package torrentv2

import "testing"

func TestFileTree(t *testing.T) {
	ft := NewFileTree()
	ft.AddFile("video.mp4", 1024*1024, "root_hash_123")
	if len(ft.Files) != 1 || ft.Files["video.mp4"].Length != 1024*1024 {
		t.Errorf("unexpected file tree: %+v", ft)
	}
}
