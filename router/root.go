package router

import (
	"Core/config"
	"Core/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	config *config.Config
	engine *gin.Engine
}

func NewRouter(config *config.Config, service *service.Service) (*Router, error) {
	r := &Router{
		config: config,
		engine: gin.New(),
	}

	r.GET("/health", r.Health)

	InitFileRouter(r, service.FileService)

	return r, r.engine.Run(config.ServerInfo.Port)
}

func (r *Router) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
