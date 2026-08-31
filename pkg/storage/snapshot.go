package storage

import "encoding/json"

type StateSnapshot struct {
	URL        string `json:"url"`
	Downloaded int64  `json:"downloaded"`
	Total      int64  `json:"total"`
}

func SerializeSnapshot(s StateSnapshot) ([]byte, error) {
	return json.Marshal(s)
}
