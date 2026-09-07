package security

import "testing"

func TestGetDefaultAuditRules(t *testing.T) {
	rules := GetDefaultAuditRules()
	if len(rules) != 2 {
		t.Errorf("expected 2 audit rules")
	}
}
