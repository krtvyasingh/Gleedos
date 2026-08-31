package subs

import "strings"

func IsTTML(content string) bool {
	return strings.Contains(content, "<tt") && strings.Contains(content, "xmlns")
}
