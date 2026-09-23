package equalizer

type EQBand struct {
	FrequencyHz float64
	GainDB      float64
}

type Preset string

const (
	VoiceClarity Preset = "voice_clarity"
	BassBoost    Preset = "bass_boost"
	Flat         Preset = "flat"
)

func GetPresetBands(p Preset) []EQBand {
	switch p {
	case VoiceClarity:
		return []EQBand{{FrequencyHz: 100, GainDB: -3}, {FrequencyHz: 2000, GainDB: 4}, {FrequencyHz: 5000, GainDB: 3}}
	case BassBoost:
		return []EQBand{{FrequencyHz: 60, GainDB: 6}, {FrequencyHz: 120, GainDB: 4}}
	default:
		return []EQBand{{FrequencyHz: 1000, GainDB: 0}}
	}
}
