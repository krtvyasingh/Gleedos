package urlutil

import (
	"net/url"
	"strings"
)

var trackingParams = map[string]bool{
	"utm_source":   true,
	"utm_medium":   true,
	"utm_campaign": true,
	"utm_term":     true,
	"utm_content":  true,
	"fbclid":       true,
	"gclid":        true,
	"si":           true,
	"feature":      true,
	"ref":          true,
}

func CleanURL(rawURL string) string {
	u, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil {
		return rawURL
	}
	q := u.Query()
	for param := range q {
		if trackingParams[strings.ToLower(param)] {
			q.Del(param)
		}
	}
	u.RawQuery = q.Encode()
	return u.String()
}
