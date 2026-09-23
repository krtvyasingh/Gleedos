package retention

import (
	"testing"
	"time"
)

func TestShouldPruneFile(t *testing.T) {
	p := RetentionPolicy{MaxAgeDays: 7}
	if !ShouldPruneFile(10*24*time.Hour, p) {
		t.Errorf("expected prune for 10-day old file")
	}
	if ShouldPruneFile(3*24*time.Hour, p) {
		t.Errorf("expected retain for 3-day old file")
	}
}
