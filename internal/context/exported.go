// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package context

import (
	"github.com/fat0troll/yapusher/internal/config"
	"github.com/rs/zerolog"
	"gitlab.com/pztrn/flagger"
)

// VERSION is the current app version
const VERSION = "0.0.1"

var (
	dlog zerolog.Logger
)

// Context is the main application context.
type Context struct {
	configFilePath string

	Config  config.Config
	Flagger *flagger.Flagger
	Logger  zerolog.Logger
}

// New creates new Context
func New() *Context {
	c := &Context{}
	return c
}
