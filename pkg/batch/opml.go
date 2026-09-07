package batch

import (
	"regexp"
	"strings"
)

var opmlURLRegex = regexp.MustCompile(`xmlUrl="([^"]+)"`)

func ParseOPMLFeeds(opmlContent string) []string {
	var urls []string
	matches := opmlURLRegex.FindAllStringSubmatch(opmlContent, -1)
	for _, m := range matches {
		if len(m) > 1 && strings.HasPrefix(m[1], "http") {
			urls = append(urls, m[1])
		}
	}
	return urls
}
