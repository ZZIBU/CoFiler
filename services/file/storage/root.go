package storage

import (
	"CoFiler/config"
	"errors"
	"io"
)

const (
	Local = "local"
)

type Storage interface {
	Save(path string, file io.Reader) error
	Delete(path string) error
	Get(path string) (io.ReadCloser, error)
}

func NewStorage(config *config.Config) (Storage, error) {
	switch config.Storage.Type {
	case Local:
		if basePath := config.Storage.BasePath; basePath == "" {
			return nil, errors.New("missing basePath for local storage")
		} else {
			return NewLocalHandler(basePath), nil
		}
	default:
		return nil, errors.New("unsupported storage type")
	}
}
