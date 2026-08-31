package subs

import "strings"

func IsSubViewerTimestamp(line string) bool {
	return strings.Count(line, ",") == 1 && strings.Count(line, ":") == 4
}
