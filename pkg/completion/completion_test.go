package completion

import (
	"strings"
	"testing"
)

func TestCompletions(t *testing.T) {
	if !strings.Contains(BashCompletion, "_gleedos_completion") {
		t.Errorf("invalid bash completion template")
	}
	if !strings.Contains(ZshCompletion, "_gleedos") {
		t.Errorf("invalid zsh completion template")
	}
}
