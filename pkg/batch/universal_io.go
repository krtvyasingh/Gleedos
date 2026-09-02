package batch

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MediaExportRecord struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Format      string `json:"format"`
	Quality     string `json:"quality"`
	CompletedAt string `json:"completed_at"`
}

func ExportPlaylist(records []MediaExportRecord, format string, outPath string) error {
	format = strings.ToLower(format)
	var content string

	switch format {
	case "json":
		data, err := json.MarshalIndent(records, "", "  ")
		if err != nil {
			return err
		}
		content = string(data)
	case "m3u", "m3u8":
		var sb strings.Builder
		sb.WriteString("#EXTM3U\n")
		for _, r := range records {
			sb.WriteString(fmt.Sprintf("#EXTINF:-1,%s\n%s\n", r.Title, r.URL))
		}
		content = sb.String()
	case "csv":
		var sb strings.Builder
		sb.WriteString("URL,Title,Format,Quality,CompletedAt\n")
		for _, r := range records {
			sb.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s\n", r.URL, r.Title, r.Format, r.Quality, r.CompletedAt))
		}
		content = sb.String()
	case "markdown", "md":
		var sb strings.Builder
		sb.WriteString("# Gleedos Exported Media Collection\n\n")
		for i, r := range records {
			sb.WriteString(fmt.Sprintf("%d. [%s](%s) - Quality: `%s`\n", i+1, r.Title, r.URL, r.Quality))
		}
		content = sb.String()
	default: // plain txt
		var sb strings.Builder
		for _, r := range records {
			sb.WriteString(r.URL + "\n")
		}
		content = sb.String()
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil && filepath.Dir(outPath) != "." {
		return err
	}
	return os.WriteFile(outPath, []byte(content), 0644)
}
