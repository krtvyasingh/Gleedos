package visionllm

type VideoSummary struct {
	Title     string
	KeyTopics []string
	Summary   string
}

func GenerateSummaryFromKeywords(keywords []string) VideoSummary {
	return VideoSummary{
		Title:     "Automated AI Overview",
		KeyTopics: keywords,
		Summary:   "Video analyzing core concepts across key segments.",
	}
}
