package dubbing

type DubTrack struct {
	LanguageCode string
	VoiceProfile string
	OutputPath   string
}

func CreateDubSpec(lang, voice, out string) DubTrack {
	return DubTrack{LanguageCode: lang, VoiceProfile: voice, OutputPath: out}
}
