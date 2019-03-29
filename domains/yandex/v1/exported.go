// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package yandexv1

import (
	"github.com/fat0troll/yapusher/internal/context"
	"github.com/rs/zerolog"
	"gitlab.com/pztrn/flagger"
)

const YANDEX_APPID = "7d8a0561fdc44c05bb6695b464403f9c"

var (
	c    *context.Context
	dlog zerolog.Logger
)

// New initializes package
func New(cc *context.Context) {
	c = cc
	dlog = c.Logger.With().Str("domain", "yandex").Int("version", 1).Logger()

	_ = c.Flagger.AddFlag(&flagger.Flag{
		Name:         "authCode",
		Description:  "Auth code obtained from Yandex (seven digits).",
		Type:         "int",
		DefaultValue: 0000000,
	})

	dlog.Info().Msg("Domain initialized")
}

// Process handles authorization and files
func Process() {
	if !checkAuth() {
		authorize()
	}
}
