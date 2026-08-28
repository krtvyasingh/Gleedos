package batch

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestExportHistoryCSV(t *testing.T) {
	tmpDir := t.TempDir()
	histPath := filepath.Join(tmpDir, "hist.json")
	store, _ := NewHistoryStore(histPath)

	_ = store.Record(HistoryRecord{
		URL:         "https://example.com/test",
		CompletedAt: time.Now(),
		OutputPath:  "/tmp/test.mp4",
		Success:     true,
	})

	csvPath := filepath.Join(tmpDir, "out.csv")
	if err := ExportHistoryCSV(store, csvPath); err != nil {
		t.Fatalf("ExportHistoryCSV failed: %v", err)
	}

	data, err := os.ReadFile(csvPath)
	if err != nil || len(data) == 0 {
		t.Errorf("empty exported CSV")
	}
}
