package equalizer

import "testing"

func TestGetPresetBands(t *testing.T) {
	bands := GetPresetBands(VoiceClarity)
	if len(bands) != 3 || bands[1].GainDB != 4 {
		t.Errorf("unexpected EQ bands: %+v", bands)
	}
}
