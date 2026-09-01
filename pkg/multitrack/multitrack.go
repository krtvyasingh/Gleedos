package multitrack

type TrackInfo struct {
	Language string
	Title    string
	IsDefault bool
}

type MediaPackage struct {
	AudioTracks []TrackInfo
	SubtitleTracks []TrackInfo
}

func NewMediaPackage() *MediaPackage {
	return &MediaPackage{}
}

func (m *MediaPackage) AddAudioTrack(lang, title string, isDefault bool) {
	m.AudioTracks = append(m.AudioTracks, TrackInfo{Language: lang, Title: title, IsDefault: isDefault})
}

func (m *MediaPackage) AddSubtitleTrack(lang, title string, isDefault bool) {
	m.SubtitleTracks = append(m.SubtitleTracks, TrackInfo{Language: lang, Title: title, IsDefault: isDefault})
}
