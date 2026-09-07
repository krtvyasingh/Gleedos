package dlna

import (
	"strings"
	"testing"
)

func TestBuildSSDPOK(t *testing.T) {
	pkt := BuildSSDPOK("http://192.168.1.10:8080/desc.xml", "uuid:gleedos-dms")
	if !strings.Contains(pkt, "HTTP/1.1 200 OK") || !strings.Contains(pkt, "uuid:gleedos-dms") {
		t.Errorf("unexpected SSDP packet: %s", pkt)
	}
}
