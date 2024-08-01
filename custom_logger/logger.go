package customlogger

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func InitLogger() (*zap.SugaredLogger, error) {
	config := zap.Config{
		Encoding:    "json", // or "console"
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder, // Lowercase level names
			EncodeTime:     zapcore.ISO8601TimeEncoder,    // Human-readable timestamps
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		return nil, err
	}

	sugar = logger.Sugar()
	return sugar, nil
}

func Sync() {
	if logger != nil {
		logger.Sync()
	}
}

// Helper function to format JSON data
func formatJSON(data []byte) string {
	var bodyMap map[string]interface{}
	if err := json.Unmarshal(data, &bodyMap); err == nil {
		prettyBody, _ := json.MarshalIndent(bodyMap, "", "    ")
		return string(prettyBody)
	}
	return string(data)
}

// Infof logs an info message with formatting and optional JSON body
func Infof(msg string, args ...interface{}) {
	if len(args) > 0 {
		if body, ok := args[len(args)-1].([]byte); ok {
			args[len(args)-1] = formatJSON(body)
		}
	}
	sugar.Infof(msg, args...)
}

// Errorf logs an error message with formatting and optional JSON body
func Errorf(msg string, args ...interface{}) {
	if len(args) > 0 {
		if body, ok := args[len(args)-1].([]byte); ok {
			args[len(args)-1] = formatJSON(body)
		}
	}
	sugar.Errorf(msg, args...)
}

// Fatalf logs a fatal message with formatting and optional JSON body
func Fatalf(msg string, args ...interface{}) {
	if len(args) > 0 {
		if body, ok := args[len(args)-1].([]byte); ok {
			args[len(args)-1] = formatJSON(body)
		}
	}
	sugar.Fatalf(msg, args...)
}

// Debugf logs a debug message with formatting and optional JSON body
func Debugf(msg string, args ...interface{}) {
	if len(args) > 0 {
		if body, ok := args[len(args)-1].([]byte); ok {
			args[len(args)-1] = formatJSON(body)
		}
	}
	sugar.Debugf(msg, args...)
}
