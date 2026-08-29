package scheduler

import (
	"testing"
	"time"
)

func TestTimeWindow(t *testing.T) {
	tw := TimeWindow{StartHour: 22, StartMin: 0, EndHour: 6, EndMin: 0}
	
	t1 := time.Date(2026, 8, 29, 23, 30, 0, 0, time.UTC)
	if !tw.IsActive(t1) {
		t.Errorf("expected active at 23:30")
	}

	t2 := time.Date(2026, 8, 29, 14, 0, 0, 0, time.UTC)
	if tw.IsActive(t2) {
		t.Errorf("expected inactive at 14:00")
	}
}
