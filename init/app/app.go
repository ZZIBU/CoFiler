package app

import (
	"CoFiler/config"
	"CoFiler/utils/logging"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

func NewServer(lc fx.Lifecycle, config *config.Config) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logging.FromContext(ctx).Info("Start to rest api server : " + config.ServerInfo.Port)
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logging.DefaultLogger().Errorw("Failed to close http server", "err", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logging.FromContext(ctx).Info("Stopped rest api server")
			return srv.Shutdown(ctx)
		},
	})
	return engine
}
