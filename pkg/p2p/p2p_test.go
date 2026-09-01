package p2p

import (
	"context"
	"testing"
)

func TestSwarmManager(t *testing.T) {
	sm := NewSwarmManager()
	if err := sm.AnnounceChunk("hash123", 9001); err != nil {
		t.Fatalf("AnnounceChunk failed: %v", err)
	}
	peers := sm.DiscoverPeers(context.Background(), "hash123")
	if len(peers) != 1 {
		t.Errorf("expected 1 peer, got %d", len(peers))
	}
}
