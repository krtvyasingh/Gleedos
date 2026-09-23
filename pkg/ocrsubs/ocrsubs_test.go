package ocrsubs

import "testing"

func TestCleanOCRText(t *testing.T) {
	cleaned := CleanOCRText(" | am ready ")
	if cleaned != "I am ready" {
		t.Errorf("unexpected OCR clean: %s", cleaned)
	}
}
