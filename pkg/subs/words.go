package subs

import "strings"

type WordTiming struct {
	Word  string
	Start float64
	End   float64
}

func ParseWordCount(text string) int {
	return len(strings.Fields(text))
}
