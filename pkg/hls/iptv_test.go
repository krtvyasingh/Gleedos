package hls

import (
	"strings"
	"testing"
)

func TestParseIPTV(t *testing.T) {
	data := `#EXTM3U
#EXTINF:-1,News HD
http://example.com/news.m3u8
`
	chans := ParseIPTV(strings.NewReader(data))
	if len(chans) != 1 || chans[0].Name != "News HD" {
		t.Errorf("unexpected IPTV channels: %+v", chans)
	}
}
