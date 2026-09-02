package stems

type AudioStems struct {
	VocalsPath string
	DrumsPath  string
	BassPath   string
	OtherPath  string
}

func FormatStemNames(baseName string) AudioStems {
	return AudioStems{
		VocalsPath: baseName + "_vocals.flac",
		DrumsPath:  baseName + "_drums.flac",
		BassPath:   baseName + "_bass.flac",
		OtherPath:  baseName + "_other.flac",
	}
}
