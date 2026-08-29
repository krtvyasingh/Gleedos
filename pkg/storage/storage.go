package storage

import (
	"context"
	"io"
)

type StorageProvider interface {
	Upload(ctx context.Context, key string, r io.Reader) error
	Name() string
}

type LocalStorage struct{}

func (l *LocalStorage) Upload(ctx context.Context, key string, r io.Reader) error {
	return nil
}

func (l *LocalStorage) Name() string {
	return "local"
}
