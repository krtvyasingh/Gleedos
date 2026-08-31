package storage

import "testing"

func TestDedupStore(t *testing.T) {
	ds := NewDedupStore()
	chunk := []byte("duplicate_stream_payload")
	if ds.IsDuplicate(chunk) {
		t.Errorf("expected first occurrence not duplicate")
	}
	if !ds.IsDuplicate(chunk) {
		t.Errorf("expected second occurrence duplicate")
	}
}
