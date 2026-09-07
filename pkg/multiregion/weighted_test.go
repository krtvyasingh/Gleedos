package multiregion

import "testing"

func TestPickWeightedRegion(t *testing.T) {
	regs := []WeightedRegion{
		{Region: "us-east-1", Weight: 10},
		{Region: "eu-west-1", Weight: 90},
	}
	pick := PickWeightedRegion(regs, 50)
	if pick != "eu-west-1" {
		t.Errorf("expected eu-west-1, got %s", pick)
	}
}
