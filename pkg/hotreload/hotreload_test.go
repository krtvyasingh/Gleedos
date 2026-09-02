package hotreload

import "testing"

func TestReloadManager(t *testing.T) {
	rm := &ReloadManager{}
	if rm.IsListening() {
		t.Errorf("expected false for nil listener")
	}
}
