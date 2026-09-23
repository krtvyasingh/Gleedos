package ocrsubs

import "strings"

func CleanOCRText(raw string) string {
	raw = strings.ReplaceAll(raw, "|", "I")
	raw = strings.TrimSpace(raw)
	return raw
}
