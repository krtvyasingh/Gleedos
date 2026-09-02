package multiregion

import "testing"

func TestHashRing(t *testing.T) {
	hr := NewHashRing([]string{"us-east-1", "eu-central-1", "ap-southeast-1"})
	r1 := hr.GetRegion("video_12345")
	if r1 == "" {
		t.Errorf("empty region selected")
	}
}
