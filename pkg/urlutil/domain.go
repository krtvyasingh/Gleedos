package urlutil

import (
	"net/url"
	"strings"
)

func ExtractDomain(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return strings.ToLower(u.Hostname())
}

func DeduplicateURLs(urls []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, u := range urls {
		clean := CleanURL(u)
		if !seen[clean] {
			seen[clean] = true
			result = append(result, clean)
		}
	}
	return result
}
