package pivoting

import "testing"

func TestPivotManager(t *testing.T) {
	pm := NewPivotManager([]string{"mirror1", "mirror2"})
	p1 := pm.PivotOnFailure()
	if p1 != "mirror2" {
		t.Errorf("expected mirror2, got %s", p1)
	}
}
