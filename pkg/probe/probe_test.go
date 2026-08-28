package probe

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInspectFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "sample.mp4")
	_ = os.WriteFile(filePath, []byte("dummy-media-content"), 0644)

	info, err := InspectFile(filePath)
	if err != nil {
		t.Fatalf("InspectFile failed: %v", err)
	}
	if info.Extension != ".mp4" || info.SizeBytes != int64(len("dummy-media-content")) {
		t.Errorf("unexpected MediaInfo: %+v", info)
	}
}
