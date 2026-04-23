# slog-helpers

```golang
package main

import (
    "flag"

    "github.com/onrik/slog-helpers"
)

func main {
    logLevel := flag.String("log-level", "info", "")
    flag.Parse()

	slog.SetDefault(slog.New(
		sloghelpers.NewTextHandler(*logLevel, nil),
	))
}

```