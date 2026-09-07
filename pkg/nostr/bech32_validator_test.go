package nostr

import "testing"

func TestBech32Validator(t *testing.T) {
	npub := "npub18005gah7g8sc2nh9ca7xmzy2nfl05wmav3nnhaqwpav5dr8n2eqsw38l2p"
	if !IsNpub(npub) {
		t.Errorf("expected valid npub")
	}
	if IsNsec(npub) {
		t.Errorf("expected false for nsec on npub")
	}
}
