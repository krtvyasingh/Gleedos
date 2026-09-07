package loudness

import "testing"

func TestGetStage1Coefficients(t *testing.T) {
	coef := GetStage1Coefficients()
	if coef.B0 <= 1.0 || coef.A2 <= 0.5 {
		t.Errorf("invalid coefficients: %+v", coef)
	}
}
