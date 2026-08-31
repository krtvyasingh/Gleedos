package storage

import "testing"

func TestNewSQLiteStore(t *testing.T) {
	s := NewSQLiteStore(":memory:")
	if s.Path != ":memory:" {
		t.Errorf("unexpected path: %s", s.Path)
	}
}
