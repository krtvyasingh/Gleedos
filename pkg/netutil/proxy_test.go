package netutil

import "testing"

func TestConfigureProxy(t *testing.T) {
	tport, err := ConfigureProxy("http://127.0.0.1:8080")
	if err != nil || tport == nil {
		t.Fatalf("ConfigureProxy failed: %v", err)
	}
}
