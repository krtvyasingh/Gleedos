package hls

import "testing"

func TestSelectBestVariant(t *testing.T) {
	variants := []VariantStream{
		{URL: "low.m3u8", Bandwidth: 500000, Resolution: "480p"},
		{URL: "high.m3u8", Bandwidth: 3000000, Resolution: "1080p"},
		{URL: "med.m3u8", Bandwidth: 1500000, Resolution: "720p"},
	}

	best := SelectBestVariant(variants)
	if best.URL != "high.m3u8" || best.Bandwidth != 3000000 {
		t.Errorf("expected high.m3u8, got %+v", best)
	}
}
