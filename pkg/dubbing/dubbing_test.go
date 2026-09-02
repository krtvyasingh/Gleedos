package dubbing

import "testing"

func TestCreateDubSpec(t *testing.T) {
	d := CreateDubSpec("es", "neural_1", "out.m4a")
	if d.LanguageCode != "es" {
		t.Errorf("unexpected dub spec: %+v", d)
	}
}
