package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger      *zap.Logger
	serviceName string
)

func InitLogger(name string) (*zap.Logger, error) {
	serviceName = name
	config := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func Sync() {
	if logger != nil {
		logger.Sync()
	}
}

func log(level, msg string, req, res interface{}) {
	fields := []zap.Field{
		zap.String("caller", serviceName),
		zap.Any("req", req),
		zap.Any("res", res),
	}

	switch level {
	case "debug":
		logger.Debug(msg, fields...)
	case "info":
		logger.Info(msg, fields...)
	case "warn":
		logger.Warn(msg, fields...)
	case "error":
		logger.Error(msg, fields...)
	case "fatal":
		logger.Fatal(msg, fields...)
	default:
		logger.Info(msg, fields...)
	}
}

func Info(msg string, req, res interface{}) {
	log("info", msg, req, res)
}

func Debug(msg string, req, res interface{}) {
	log("debug", msg, req, res)
}

func Warn(msg string, req, res interface{}) {
	log("warn", msg, req, res)
}

func Error(msg string, req, res interface{}) {
	log("error", msg, req, res)
}

func Fatal(msg string, req, res interface{}) {
	log("fatal", msg, req, res)
}
