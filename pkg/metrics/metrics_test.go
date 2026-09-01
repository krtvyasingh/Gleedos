package metrics

import (
	"strings"
	"testing"
)

func TestGetDashboardTemplate(t *testing.T) {
	json := GetDashboardTemplate()
	if !strings.Contains(json, "Gleedos High-Performance Engine Telemetry") {
		t.Errorf("unexpected dashboard JSON: %s", json)
	}
}
