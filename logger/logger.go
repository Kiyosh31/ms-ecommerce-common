package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger      *zap.Logger
	Sugar       *zap.SugaredLogger
	serviceName string
)

func InitLogger(name string) {
	serviceName = name
	config := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:  "timestamp",
			LevelKey: "level",
			NameKey:  "logger",
			// CallerKey:      "caller", // Remove this line
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
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	Sugar = Logger.Sugar()
}

func Sync() {
	if Logger != nil {
		Logger.Sync()
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
		Logger.Debug(msg, fields...)
	case "info":
		Logger.Info(msg, fields...)
	case "warn":
		Logger.Warn(msg, fields...)
	case "error":
		Logger.Error(msg, fields...)
	case "fatal":
		Logger.Fatal(msg, fields...)
	default:
		Logger.Info(msg, fields...)
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
