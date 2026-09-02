package ghostsuggest

import "strings"

func SuggestNext(input string, history []string) string {
	for _, h := range history {
		if strings.HasPrefix(h, input) && h != input {
			return h
		}
	}
	return ""
}
