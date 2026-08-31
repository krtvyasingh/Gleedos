package subs

import "strings"

func IsMicroDVD(line string) bool {
	return strings.HasPrefix(line, "{") && strings.Contains(line, "}{")
}
