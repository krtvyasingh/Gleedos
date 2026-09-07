package dlna

import (
	"strings"
	"testing"
)

func TestFormatDIDLItem(t *testing.T) {
	didl := FormatDIDLItem("1", "0", "Video", "http://127.0.0.1/1.mp4", "video/mp4")
	if !strings.Contains(didl, `<dc:title>Video</dc:title>`) {
		t.Errorf("unexpected DIDL item: %s", didl)
	}
}
