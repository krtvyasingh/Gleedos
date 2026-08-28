package signals

import "testing"

func TestSetupSignalHandler(t *testing.T) {
	ctx, cancel := SetupSignalHandler()
	if ctx == nil || cancel == nil {
		t.Fatal("expected non-nil context and cancel func")
	}
	cancel()
}
