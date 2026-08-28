package batch

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type JSONPlaylist struct {
	Name string   `json:"name"`
	URLs []string `json:"urls"`
}

func ParsePlaylistFile(path string) ([]string, error) {
	ext := strings.ToLower(filepath.Ext(path))
	if ext == ".json" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		var pl JSONPlaylist
		if err := json.Unmarshal(data, &pl); err == nil && len(pl.URLs) > 0 {
			return pl.URLs, nil
		}
		var list []string
		if err := json.Unmarshal(data, &list); err == nil {
			return list, nil
		}
	}
	return ParseURLFile(path)
}
