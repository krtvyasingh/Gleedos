package dash

import (
	"strings"
	"testing"
)

func TestParseMPD(t *testing.T) {
	xmlData := `<?xml version="1.0"?>
<MPD type="static">
  <Period>
    <AdaptationSet mimeType="video/mp4">
      <Representation id="1080p" bandwidth="5000000" width="1920" height="1080"/>
    </AdaptationSet>
  </Period>
</MPD>`

	mpd, err := ParseMPD(strings.NewReader(xmlData))
	if err != nil || len(mpd.Period) == 0 {
		t.Fatalf("ParseMPD failed: %v", err)
	}
	if mpd.Period[0].AdaptationSet[0].Representation[0].Height != 1080 {
		t.Errorf("unexpected height in representation")
	}
}
