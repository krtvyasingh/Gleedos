package multitrack

import "testing"

func TestMediaPackage(t *testing.T) {
	pkg := NewMediaPackage()
	pkg.AddAudioTrack("eng", "English Audio", true)
	pkg.AddSubtitleTrack("spa", "Spanish Subs", false)
	if len(pkg.AudioTracks) != 1 || len(pkg.SubtitleTracks) != 1 {
		t.Errorf("unexpected tracks count: %+v", pkg)
	}
}
