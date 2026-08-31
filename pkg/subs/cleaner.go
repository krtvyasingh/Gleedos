package subs

import "regexp"

var tagReg = regexp.MustCompile(`<[^>]*>`)

func StripHTMLTags(s string) string {
	return tagReg.ReplaceAllString(s, "")
}
