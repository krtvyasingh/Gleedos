package batch

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExportPlaylistFormats(t *testing.T) {
	tmpDir := t.TempDir()
	records := []MediaExportRecord{
		{URL: "https://example.com/1", Title: "Track 1", Format: "mp4", Quality: "1080p"},
		{URL: "https://example.com/2", Title: "Track 2", Format: "mp3", Quality: "audio"},
	}

	for _, fmtName := range []string{"json", "csv", "m3u8", "md", "txt"} {
		outPath := filepath.Join(tmpDir, "export."+fmtName)
		if err := ExportPlaylist(records, fmtName, outPath); err != nil {
			t.Fatalf("ExportPlaylist failed for %s: %v", fmtName, err)
		}
		data, err := os.ReadFile(outPath)
		if err != nil || len(data) == 0 {
			t.Errorf("empty export file for format %s", fmtName)
		}
	}
}
