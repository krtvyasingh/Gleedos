package security

import "testing"

func TestValidateArgument(t *testing.T) {
	if err := ValidateArgument("valid_param"); err != nil {
		t.Errorf("expected valid: %v", err)
	}
	if err := ValidateArgument("param; rm -rf /"); err == nil {
		t.Errorf("expected error on command chaining")
	}
}
