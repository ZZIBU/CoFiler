package router

import (
	"Core/config"
	"Core/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	config  *config.Config
	engine  *gin.Engine
	service *service.Service
}

func NewRouter(config *config.Config, service *service.Service) (*Router, error) {
	r := &Router{
		config:  config,
		engine:  gin.New(),
		service: service,
	}

	//r.engine.Use(requestTimeOutMiddleWare(5 * time.Second))
	//
	//NewMongoRouter(r, r.service.MService)

	return r, r.engine.Run(config.ServerInfo.Port)
}
