package service

import (
	"Core/config"
	"Core/repository"
	"Core/storage"
)

type Service struct {
	config      *config.Config
	FileService *FileService
}

func NewService(
	config *config.Config,
	repository *repository.Repository,
	storage storage.Storage) (*Service, error) {
	r := &Service{
		config:      config,
		FileService: NewFileService(storage),
	}

	return r, nil
}
