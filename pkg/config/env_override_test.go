package config

import "testing"

func TestResolveConfigString(t *testing.T) {
	s := ResolveConfigString("NON_EXISTENT_GLEEDOS_VAR", "default_val")
	if s != "default_val" {
		t.Errorf("expected default value")
	}
}
