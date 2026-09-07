package tlscloak

import "testing"

func TestValidateALPN(t *testing.T) {
	if !ValidateALPN([]string{"h2", "http/1.1"}) {
		t.Errorf("expected valid ALPN")
	}
	if ValidateALPN([]string{"ftp"}) {
		t.Errorf("expected invalid ALPN")
	}
}
