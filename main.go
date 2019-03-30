// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package main

import (
	"fmt"
	"github.com/fat0troll/yapusher/domains/yandex/v1"
	"github.com/fat0troll/yapusher/internal/context"
	"runtime/debug"
)

// In production builds these variables are set by goreleaser
var (
	version = "master"
	commit  = "?"
	date    = ""
)

//nolint:gochecknoinits
func init() {
	if info, available := debug.ReadBuildInfo(); available {
		if date == "" && info.Main.Version != "(devel)" {
			version = info.Main.Version
			commit = fmt.Sprintf("(unknown, mod sum: %q)", info.Main.Sum)
			date = "(unknown)"
		}
	}
}

func main() {
	c := context.New()
	c.Init()
	c.Logger.Info().Str("version", version).Str("commit", commit).Str("build date", date).Msg("yapusher is starting")
	c.InitConfig()

	// Initializing domains
	yandexv1.New(c)

	// Parsing app flags
	c.Flagger.Parse()

	// Authorizing to Yandex
	yandexv1.Process()
}
