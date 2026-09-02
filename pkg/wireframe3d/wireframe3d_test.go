package wireframe3d

import (
	"strings"
	"testing"
)

func TestRender3DCube(t *testing.T) {
	cube := Render3DCube(45)
	if !strings.Contains(cube, "+------+") {
		t.Errorf("unexpected wireframe: %s", cube)
	}
}
