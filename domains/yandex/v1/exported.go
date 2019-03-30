// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package yandexv1

import (
	"github.com/fat0troll/yapusher/internal/context"
	"github.com/rs/zerolog"
	"gitlab.com/pztrn/flagger"
)

const YANDEX_APPID = "7d8a0561fdc44c05bb6695b464403f9c"
const YANDEX_APPPW = "56e12e4ed0d64738bf441a47f68c7146"
const DEVICE_NAME = "yapusher-cli"

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

	_ = c.Flagger.AddFlag(&flagger.Flag{
		Name:         "uploadPath",
		Description:  "Path to upload your file on Yandex.Disk. Must exist before uploading.",
		Type:         "string",
		DefaultValue: "/",
	})

	_ = c.Flagger.AddFlag(&flagger.Flag{
		Name:         "file",
		Description:  "Path to file that will be uploaded. Max upload size - 50 GB",
		Type:         "string",
		DefaultValue: "",
	})

	dlog.Info().Msg("Domain initialized")
}

// Process handles authorization and files
func Process() {
	authCode, _ := c.Flagger.GetIntValue("authCode")
	if authCode != 0 {
		sendCode(authCode)
	}

	filePath, _ := c.Flagger.GetStringValue("file")
	if filePath != "" {
		uploadFile()
	}

	if !checkAuth() {
		authorize()
	}

	showHelp()
}
