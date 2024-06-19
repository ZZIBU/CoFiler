package app

import (
	"Core/config"
	"Core/repository"
	"Core/router"
	"Core/service"
	"Core/storage"
)

type App struct {
	config     *config.Config
	router     *router.Router
	service    *service.Service
	repository *repository.Repository
	storage    storage.Storage
}

func NewApp(config *config.Config) {
	app := &App{
		config: config,
	}

	var err error
	if app.storage, err = storage.NewStorage(config); err != nil {
		panic(err)
	}
	if app.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}
	if app.service, err = service.NewService(config, app.repository, app.storage); err != nil {
		panic(err)
	}
	if app.router, err = router.NewRouter(config, app.service); err != nil {
		panic(err)
	}
}
