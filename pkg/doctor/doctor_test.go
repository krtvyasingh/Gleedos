package doctor

import "testing"

func TestRunDiagnostics(t *testing.T) {
	results := RunDiagnostics()
	if len(results) != 3 {
		t.Errorf("expected 3 tool checks, got %d", len(results))
	}
}
