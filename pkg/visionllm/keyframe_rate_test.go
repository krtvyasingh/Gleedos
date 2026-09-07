package visionllm

import (
	"testing"
	"time"
)

func TestCalculateKeyframeInterval(t *testing.T) {
	iv := CalculateKeyframeInterval(100*time.Second, 10)
	if iv != 10*time.Second {
		t.Errorf("expected 10s, got %v", iv)
	}
}
