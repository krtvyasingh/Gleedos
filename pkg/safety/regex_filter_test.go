package safety

import "testing"

func TestDetectSecrets(t *testing.T) {
	if !DetectSecrets(`api_key = "secret_value_123"`) {
		t.Errorf("expected secret detection")
	}
	if DetectSecrets("Just regular plain text") {
		t.Errorf("expected no secret detection")
	}
}
