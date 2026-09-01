package xattr

type MediaMetadata struct {
	OriginalURL string
	Author      string
	Title       string
}

func FormatXAttrMap(meta MediaMetadata) map[string]string {
	return map[string]string{
		"com.apple.metadata:kMDItemWhereFroms": meta.OriginalURL,
		"com.apple.metadata:kMDItemAuthors":    meta.Author,
		"com.apple.metadata:kMDItemTitle":      meta.Title,
	}
}
