package tagger

type ThumbnailInfo struct {
	MimeType string
	Data     []byte
}

func NewThumbnail(mime string, data []byte) ThumbnailInfo {
	return ThumbnailInfo{MimeType: mime, Data: data}
}
