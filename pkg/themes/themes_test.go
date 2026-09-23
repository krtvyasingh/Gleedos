package themes

import "testing"

func TestGetTheme(t *testing.T) {
	th := GetTheme("nord")
	if th.Accent != "#88C0D0" {
		t.Errorf("unexpected accent: %s", th.Accent)
	}
}
