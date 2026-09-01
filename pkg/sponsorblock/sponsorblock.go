package sponsorblock

type SegmentType string

const (
	Sponsor      SegmentType = "sponsor"
	Intro        SegmentType = "intro"
	Outro        SegmentType = "outro"
	SelfPromo    SegmentType = "selfpromo"
)

type SkipSegment struct {
	Category  SegmentType
	StartTime float64
	EndTime   float64
}

func ShouldSkipTime(segments []SkipSegment, timestamp float64) bool {
	for _, s := range segments {
		if timestamp >= s.StartTime && timestamp <= s.EndTime {
			return true
		}
	}
	return false
}
