package mouse

import "testing"

func TestParseSequence(t *testing.T) {
	ev := ParseSequence("\x1b[M")
	if ev.Action != Click {
		t.Errorf("unexpected action: %v", ev)
	}
}
