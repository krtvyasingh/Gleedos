package ghostsuggest

import "testing"

func TestSuggestNext(t *testing.T) {
	hist := []string{"gleedos --audio https://music.com", "gleedos doctor"}
	sug := SuggestNext("gleedos --a", hist)
	if sug != "gleedos --audio https://music.com" {
		t.Errorf("unexpected suggestion: %s", sug)
	}
}
