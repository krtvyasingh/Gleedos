package keychain

import "testing"

func TestParseNetscapeCookies(t *testing.T) {
	cookies := ParseNetscapeCookies([]string{"val1", "val2"})
	if len(cookies) != 2 {
		t.Errorf("expected 2 cookies")
	}
}
