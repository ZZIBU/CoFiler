package controller

import (
	"CoFiler/utils/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MetricRouter(e *gin.Engine) {
	metricV1 := e.Group("/api/v1")
	metricV1.Use()
	{
		metricV1.GET("health", Health)
	}
}

func Health(c *gin.Context) {
	logger := logging.FromContext(c)
	logger.Info("HealthCheck")

	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
