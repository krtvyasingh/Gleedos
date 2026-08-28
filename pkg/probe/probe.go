package probe

import (
	"fmt"
	"os"
	"path/filepath"
)

type MediaInfo struct {
	Path      string
	Extension string
	SizeBytes int64
}

func InspectFile(path string) (*MediaInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("stat failed: %w", err)
	}
	return &MediaInfo{
		Path:      path,
		Extension: filepath.Ext(path),
		SizeBytes: info.Size(),
	}, nil
}
