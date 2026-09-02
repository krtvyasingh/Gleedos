package k8s

import (
	"strings"
	"testing"
)

func TestGenerateCRDYAML(t *testing.T) {
	yaml := GenerateCRDYAML("job-1", "https://example.com/v.mp4")
	if !strings.Contains(yaml, "GleedosDownloadJob") {
		t.Errorf("unexpected YAML output")
	}
}
