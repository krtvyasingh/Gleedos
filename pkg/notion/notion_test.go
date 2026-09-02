package notion

import (
	"strings"
	"testing"
)

func TestExportToObsidianMarkdown(t *testing.T) {
	md := ExportToObsidianMarkdown(NoteEntry{Title: "Video Note", URL: "https://example.com", Duration: "10:00", Summary: "Cool video"})
	if !strings.Contains(md, "# Video Note") || !strings.Contains(md, "Cool video") {
		t.Errorf("unexpected markdown: %s", md)
	}
}
