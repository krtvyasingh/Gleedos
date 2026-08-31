package storage

import "testing"

func TestSerializeSnapshot(t *testing.T) {
	data, err := SerializeSnapshot(StateSnapshot{URL: "https://example.com", Downloaded: 100, Total: 200})
	if err != nil || len(data) == 0 {
		t.Fatalf("SerializeSnapshot failed: %v", err)
	}
}
