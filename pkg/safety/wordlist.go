package safety

import "strings"

type SafetyFilter struct {
	bannedTerms map[string]struct{}
}

func NewSafetyFilter(terms []string) *SafetyFilter {
	m := make(map[string]struct{})
	for _, t := range terms {
		m[strings.ToLower(t)] = struct{}{}
	}
	return &SafetyFilter{bannedTerms: m}
}

func (s *SafetyFilter) ContainsBanned(text string) bool {
	words := strings.Fields(strings.ToLower(text))
	for _, w := range words {
		if _, ok := s.bannedTerms[w]; ok {
			return true
		}
	}
	return false
}
