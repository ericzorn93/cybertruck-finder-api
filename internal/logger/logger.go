package logger

import (
	"log/slog"
	"os"
)

// Logger is a global logger instance that can be used throughout the application.
// It is initialized with default options and can be customized as needed.
var Logger *slog.Logger

func init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug, // Set minimum level to debug
	}

	Logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
