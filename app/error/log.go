package error

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ApiLogger is a custom logger for the API
type ApiLogger struct {
	*slog.Logger
}

// HandleValues is a custom log handler for the echo logger middleware
// It handles the log values and logs them into the slog logger sink
func (l ApiLogger) HandleValues(c echo.Context, v middleware.RequestLoggerValues) error {
	if v.Error == nil {
		l.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST", slog.String("uri", v.URI), slog.Int("status", v.Status))
	} else {
		l.LogAttrs(context.Background(), slog.LevelError, "ERROR", slog.String("uri", v.URI), slog.Int("status", v.Status), slog.String("error", v.Error.Error()))
	}
	// This should not return an error
	return nil
}

// Create an error message that is fatal and exists the process
func (l ApiLogger) Fatal(msg string) {
	l.LogAttrs(context.Background(), slog.LevelError, "FATAL", slog.String("msg", msg))
	os.Exit(1)
}

// Returns a new ApiLogger with the given log level
func NewApiLogger(level string) *ApiLogger {
	return &ApiLogger{Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: ToLogLevel(level)}))}
}

// Returns the corresponding slog log level for the given string
func ToLogLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
