package repository

import (
	"Core/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	config *config.Config
	client *mongo.Client
	db     *mongo.Database

	// TODO: 사용할 컬렉션
}

func NewRepository(config *config.Config) (*Repository, error) {
	repo := &Repository{
		config: config,
	}

	return repo, nil
}
