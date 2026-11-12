package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func New(config *Config) *Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(config.LogLevel),
	}

	stdoutHandler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(stdoutHandler)
	logger.With()

	return &Logger{
		Logger: logger,
	}
}
