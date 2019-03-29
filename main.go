// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package main

import (
	"github.com/fat0troll/yapusher/domains/yandex/v1"
	"github.com/fat0troll/yapusher/internal/context"
)

func main() {
	c := context.New()
	c.Init()
	c.InitConfig()

	// Initializing domains
	yandexv1.New(c)

	// Parsing app flags
	c.Flagger.Parse()

	// Authorizing to Yandex
	yandexv1.Process()
}
