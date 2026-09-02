package notion

import "fmt"

type NoteEntry struct {
	Title     string
	URL       string
	Summary   string
	Duration  string
}

func ExportToObsidianMarkdown(n NoteEntry) string {
	return fmt.Sprintf(`# %s

- **Source URL**: %s
- **Duration**: %s

## Summary
%s
`, n.Title, n.URL, n.Duration, n.Summary)
}
