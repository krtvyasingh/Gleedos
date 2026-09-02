package safety

type ContentRating string

const (
	Safe     ContentRating = "G"
	Moderate ContentRating = "PG-13"
	Explicit ContentRating = "R"
)

func ClassifyContent(flagScore float64) ContentRating {
	if flagScore > 0.8 {
		return Explicit
	}
	if flagScore > 0.3 {
		return Moderate
	}
	return Safe
}
