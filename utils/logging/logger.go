package logging

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

const loggerKey = "logger"

var (
	defaultLogger     *zap.SugaredLogger
	defaultLoggerOnce sync.Once
	conf              = &LoggerConfig{
		Encoding:    "console",
		Level:       zapcore.InfoLevel,
		Development: true,
	}
)

type LoggerConfig struct {
	Encoding    string
	Level       zapcore.Level
	Development bool
}

func SetConfig(c *LoggerConfig) {
	conf = &LoggerConfig{
		Encoding:    c.Encoding,
		Level:       c.Level,
		Development: c.Development,
	}
}

func SetLevel(l zapcore.Level) {
	conf.Level = l
}

func NewLogger(conf *LoggerConfig) *zap.SugaredLogger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg := zap.Config{
		Encoding:         conf.Encoding,
		EncoderConfig:    ec,
		Level:            zap.NewAtomicLevelAt(conf.Level),
		Development:      conf.Development,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	return logger.Sugar()
}

func DefaultLogger() *zap.SugaredLogger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewLogger(conf)
	})
	return defaultLogger
}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	if gCtx, ok := ctx.(*gin.Context); ok {
		ctx = gCtx.Request.Context()
	}
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return DefaultLogger()
	}
	if gCtx, ok := ctx.(*gin.Context); ok && gCtx != nil {
		ctx = gCtx.Request.Context()
	}
	if logger, ok := ctx.Value(loggerKey).(*zap.SugaredLogger); ok {
		return logger
	}
	return DefaultLogger()
}
