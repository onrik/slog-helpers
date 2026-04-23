package sloghelpers

import (
	"fmt"
	"log/slog"
	"reflect"
	"strings"
)

var (
	stringerType = reflect.TypeFor[fmt.Stringer]()
)

func ReplaceAttr(groups []string, a slog.Attr) slog.Attr {
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
		return slog.String(a.Key, FormatValue(a.Value.Any()))
	}

	return a
}

func FormatValue(object any) string {
	v := reflect.ValueOf(object)
	if !v.IsValid() || v.Type().Implements(stringerType) {
		return fmt.Sprintf("%+v", object)
	}

	if v.Type().Kind() == reflect.Pointer && !v.IsNil() {
		return FormatValue(v.Elem().Interface())
	}

	return fmt.Sprintf("%+v", object)
}
