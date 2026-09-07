package wireframe3d

import "testing"

func TestProjectPoint(t *testing.T) {
	p := Point3D{X: 1.0, Y: 1.0, Z: 0.0}
	proj := ProjectPoint(p, 0.0, 10.0)
	if proj.X != 2 || proj.Y != 2 {
		t.Errorf("unexpected projection: %+v", proj)
	}
}
