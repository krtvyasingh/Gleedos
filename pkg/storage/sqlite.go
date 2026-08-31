package storage

type CacheEntry struct {
	Key       string
	Value     []byte
	ExpiresAt int64
}

type SQLiteStore struct {
	Path string
}

func NewSQLiteStore(path string) *SQLiteStore {
	return &SQLiteStore{Path: path}
}
