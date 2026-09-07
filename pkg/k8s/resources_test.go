package k8s

import (
	"strings"
	"testing"
)

func TestFormatResourceRequirements(t *testing.T) {
	res := FormatResourceRequirements("2", "4Gi")
	if !strings.Contains(res, "cpu: 2") || !strings.Contains(res, "memory: 4Gi") {
		t.Errorf("unexpected resource yaml: %s", res)
	}
}
