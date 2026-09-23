package turnstilehook

import "testing"

func TestSolverEvent(t *testing.T) {
	ev := NewSolverEvent("key123", "login")
	if ev.SiteKey != "key123" {
		t.Errorf("unexpected event: %+v", ev)
	}
}
