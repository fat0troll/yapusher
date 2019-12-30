// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package yandexv1

import (
	"os"
)

func showHelp() {
	dlog.Info().Msg("This app is authorized for uploading your files one by one to Yandex.Disk.")
	dlog.Info().Msg("For information how to use this app, run yapusher with -h flag or head to https://source.hodakov.me/fat0troll/yapusher/blob/master/README.")

	os.Exit(0)
}
