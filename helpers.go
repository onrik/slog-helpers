package sloghelpers

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strings"
)

var (
	stringerType = reflect.TypeOf((*Stringer)(nil)).Elem()
)

type Stringer interface {
	String() string
}

type Handler struct {
	slog.Handler
}

func NewTextHandler(level string) slog.Handler {
	return &Handler{
		slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			AddSource:   true,
			Level:       parseLevel(level),
			ReplaceAttr: replaceAttr,
		}),
	}
}

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.SourceKey {
		if s, ok := a.Value.Any().(*slog.Source); ok {
			parts := strings.Split(s.File, "/")
			var f string
			if len(parts) >= 4 {
				f = strings.Join(parts[len(parts)-2:], "/")
			} else {
				f = strings.Join(parts, "/")
			}
			return slog.String(
				a.Key,
				fmt.Sprintf("%s:%d", f, s.Line),
			)
		}
	}

	if a.Key == "error" {
		return a
	}

	if a.Value.Kind() == slog.KindAny {
		return slog.String(a.Key, formatValue(a.Value.Any()))
	}

	return a
}

// formatValue - разыменовывает непустые указатели
func formatValue(object any) string {
	v := reflect.ValueOf(object)
	if !v.IsValid() || v.Type().Implements(stringerType) {
		return fmt.Sprintf("%+v", object)
	}

	if v.Type().Kind() == reflect.Pointer && !v.IsNil() {
		return formatValue(v.Elem().Interface())
	}

	return fmt.Sprintf("%+v", object)
}

func parseLevel(level string) slog.Level {
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

	slog.With("level", level).Warn("Invalid log level")
	return slog.LevelInfo
}
