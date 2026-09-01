package mobile

import "strings"

func ParseMobileSharePayload(rawText string) string {
	lines := strings.Fields(rawText)
	for _, l := range lines {
		if strings.HasPrefix(l, "http://") || strings.HasPrefix(l, "https://") {
			return l
		}
	}
	return ""
}
