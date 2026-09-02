package dlna

import (
	"strings"
	"testing"
)

func TestGetDeviceXML(t *testing.T) {
	xml := GetDeviceXML()
	if !strings.Contains(xml, "Gleedos DLNA Server") {
		t.Errorf("unexpected device XML: %s", xml)
	}
}
