package deploy

import (
	"strings"
	"testing"
)

func TestGetDistrolessDockerfile(t *testing.T) {
	df := GetDistrolessDockerfile()
	if !strings.Contains(df, "distroless") {
		t.Errorf("unexpected Dockerfile: %s", df)
	}
}
