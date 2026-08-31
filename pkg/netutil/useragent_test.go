package netutil

import "testing"

func TestGetUserAgent(t *testing.T) {
	ua := GetUserAgent(0)
	if ua == "" {
		t.Errorf("empty user agent")
	}
}
