package sloghelpers

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

type Handler struct {
	slog.Handler
}

func NewTextHandler(level string, writer io.Writer) *Handler {
	if writer == nil {
		writer = os.Stderr
	}

	return &Handler{
		slog.NewTextHandler(writer, &slog.HandlerOptions{
			AddSource:   true,
			Level:       ParseLevel(level),
			ReplaceAttr: ReplaceAttr,
		}),
	}
}

func ParseLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case slog.LevelError.String():
		return slog.LevelError
	case slog.LevelWarn.String():
		return slog.LevelWarn
	case slog.LevelInfo.String():
		return slog.LevelInfo
	case slog.LevelDebug.String():
		return slog.LevelDebug
	}

	slog.With("level", level).Warn("Slog level invalid")
	return slog.LevelInfo
}
