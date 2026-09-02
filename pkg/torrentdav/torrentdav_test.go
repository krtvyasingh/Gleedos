package torrentdav

import "testing"

func TestMagnetGateway(t *testing.T) {
	gw := &MagnetGateway{Port: 8080}
	url := gw.GetStreamURL("abcd1234efgh")
	if url != "http://127.0.0.1:8080/stream/abcd1234efgh" {
		t.Errorf("unexpected stream URL: %s", url)
	}
}
