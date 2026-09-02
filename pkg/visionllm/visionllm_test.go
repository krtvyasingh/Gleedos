package visionllm

import "testing"

func TestGenerateSummaryFromKeywords(t *testing.T) {
	s := GenerateSummaryFromKeywords([]string{"AI", "Space", "Physics"})
	if len(s.KeyTopics) != 3 {
		t.Errorf("unexpected topics count: %+v", s)
	}
}
