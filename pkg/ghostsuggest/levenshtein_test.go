package ghostsuggest

import "testing"

func TestLevenshtein(t *testing.T) {
	d := Levenshtein("gleedos", "gleedos-cli")
	if d != 4 {
		t.Errorf("expected distance 4, got %d", d)
	}
}
