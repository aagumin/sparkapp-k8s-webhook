package webhook

import (
	"log"
	"log/slog"
	"os"
	"strings"
)

func InitLogger() *log.Logger {
	var logLevel slog.Level
	logLevelValue := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	switch logLevelValue {
	case "":
		logLevel = slog.LevelInfo
	case "INFO":
		logLevel = slog.LevelInfo
	case "DEBUG":
		logLevel = slog.LevelDebug
	case "ERROR":
		logLevel = slog.LevelError
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), logLevel)
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, opts)))
	return logger
}
