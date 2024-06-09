package service

import (
	"Core/config"
	"Core/repository"
)

type Service struct {
	config     *config.Config
	repository *repository.Repository
}

func NewService(config *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{
		config:     config,
		repository: repository,
	}

	return r, nil
}
