package ipfs

import (
	"strings"
	"testing"
)

func TestGenerateMockCID(t *testing.T) {
	cid := GenerateMockCID([]byte("ipfs_payload"))
	if !strings.HasPrefix(cid, "Qm") {
		t.Errorf("unexpected CID: %s", cid)
	}
}
