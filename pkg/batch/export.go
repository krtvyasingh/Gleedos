package batch

import (
	"encoding/csv"
	"os"
	"time"
)

func ExportHistoryCSV(store *HistoryStore, outPath string) error {
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	_ = w.Write([]string{"URL", "CompletedAt", "OutputPath", "Success", "Error"})
	for _, r := range store.GetAll() {
		_ = w.Write([]string{
			r.URL,
			r.CompletedAt.Format(time.RFC3339),
			r.OutputPath,
			map[bool]string{true: "true", false: "false"}[r.Success],
			r.ErrorMessage,
		})
	}
	return nil
}
