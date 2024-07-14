package main

import (
	"CoFiler/config"
	"CoFiler/init/app"
	"CoFiler/rpc"
	"CoFiler/services/file"
	"CoFiler/services/file/storage"
	"CoFiler/services/metric"
	"CoFiler/utils/logging"
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var envFlag = flag.String("config", "./env.toml", "env file not found")

func printAppInfo(config *config.Config) {
	b, _ := json.MarshalIndent(&config, "", "  ")
	logging.DefaultLogger().Infof("application information\n%s", string(b))
}

func main() {
	flag.Parse()
	conf := config.NewConfig(*envFlag)
	logging.SetConfig(&logging.LoggerConfig{
		Encoding:    conf.Logging.Encoding,
		Level:       zapcore.Level(conf.Logging.Level),
		Development: conf.ServerInfo.Development,
	})
	defer logging.DefaultLogger().Sync()

	// DI + Start Gin
	fx.New(
		fx.Supply(conf),
		fx.Supply(logging.DefaultLogger().Desugar()),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log.Named("fx")}
		}),
		fx.Invoke(
			printAppInfo,
		),
		fx.Provide(
			storage.NewStorage,
			file.NewService,
			file.NewHandler,
			rpc.NewCofilerClient,

			// gin di
			app.NewServer,
		),
		fx.Invoke(
			file.NewRouter,
			metric.NewRouter,
			func(*gin.Engine) {},
		),
	).Run()
}
