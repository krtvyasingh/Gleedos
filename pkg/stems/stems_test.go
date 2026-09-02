package stems

import "testing"

func TestFormatStemNames(t *testing.T) {
	stems := FormatStemNames("track")
	if stems.VocalsPath != "track_vocals.flac" {
		t.Errorf("unexpected stems: %+v", stems)
	}
}
