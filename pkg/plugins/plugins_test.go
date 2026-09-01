package plugins

import "testing"

func TestFilterPlugin(t *testing.T) {
	p := &FilterPlugin{Name: "cleaner", Version: "1.0.0"}
	if p.TransformURL("https://example.com") != "https://example.com" {
		t.Errorf("unexpected transform output")
	}
}
