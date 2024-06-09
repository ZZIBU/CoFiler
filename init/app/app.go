package app

import (
	"Core/config"
	"Core/repository"
	"Core/router"
	"Core/service"
)

type App struct {
	config     *config.Config
	router     *router.Router
	service    *service.Service
	repository *repository.Repository
}

func NewApp(config *config.Config) {
	app := &App{
		config: config,
	}

	var err error
	if app.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}
	if app.service, err = service.NewService(config, app.repository); err != nil {
		panic(err)
	}
	if app.router, err = router.NewRouter(config, app.service); err != nil {
		panic(err)
	}
}
