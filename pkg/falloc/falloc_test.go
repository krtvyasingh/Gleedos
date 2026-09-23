package falloc

import "testing"

func TestPlanPreallocation(t *testing.T) {
	plan := PlanPreallocation(1024 * 1024)
	if plan.TotalBytes != 1024*1024 || plan.BlockSize != 4096 {
		t.Errorf("unexpected plan: %+v", plan)
	}
}
