package voicecmd

import "strings"

func MatchHotword(transcript string) (bool, string) {
	lower := strings.ToLower(transcript)
	if strings.Contains(lower, "gleedos download") {
		parts := strings.Split(lower, "gleedos download")
		if len(parts) > 1 {
			return true, strings.TrimSpace(parts[1])
		}
	}
	return false, ""
}
