package security

import "testing"

func TestRedactSecrets(t *testing.T) {
	redacted := RedactSecrets("https://example.com/api?token=supersecret123&user=admin")
	if redacted != "https://example.com/api?token=REDACTED&user=admin" {
		t.Errorf("unexpected redacted output: %s", redacted)
	}
}
