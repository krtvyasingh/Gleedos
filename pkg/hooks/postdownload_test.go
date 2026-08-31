package hooks

import "testing"

func TestExecuteHook(t *testing.T) {
	if err := ExecuteHook("", "/tmp/test.mp4"); err != nil {
		t.Errorf("expected nil error on empty hook")
	}
}
