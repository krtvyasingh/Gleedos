package airplay

import (
	"strings"
	"testing"
)

func TestFormatPlaybackInfoPlist(t *testing.T) {
	plist := FormatPlaybackInfoPlist(120.0, 30.0)
	if !strings.Contains(plist, "<key>duration</key>") {
		t.Errorf("unexpected plist: %s", plist)
	}
}
