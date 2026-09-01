package nfo

import (
	"strings"
	"testing"
)

func TestGenerateNFOXML(t *testing.T) {
	xml := GenerateNFOXML(MovieNFO{Title: "Inception", Year: 2010, Plot: "Dreams"})
	if !strings.Contains(xml, "<title>Inception</title>") {
		t.Errorf("unexpected NFO XML: %s", xml)
	}
}
