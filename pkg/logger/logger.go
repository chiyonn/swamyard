package logger

import (
	"fmt"
	"log/slog"
	"os"
)

var log *slog.Logger

// InitLogger configures the global structured logger using slog.
func InitLogger() {
	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

// Info logs an informational message using fmt-style formatting.
func Info(msg string, args ...interface{}) {
	log.Info(fmt.Sprintf(msg, args...))
}

// Error logs an error message using fmt-style formatting.
func Error(msg string, args ...interface{}) {
	log.Error(fmt.Sprintf(msg, args...))
}
