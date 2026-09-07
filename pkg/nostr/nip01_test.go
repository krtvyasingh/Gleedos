package nostr

import "testing"

func TestComputeEventID(t *testing.T) {
	id := ComputeEventID("pubkey123", "content", 1700000000, 1)
	if len(id) != 64 {
		t.Errorf("invalid event ID: %s", id)
	}
}
