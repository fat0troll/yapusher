// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package yandexv1

import (
	"os"
	"strings"
)

// authorize autorizes user and saves token to config
// Autorization made of two steps: providing user URL to create token and
// getting token from provided code
func authorize() {
	if !checkAuth() {
		baseURL := "https://oauth.yandex.ru/authorize?response_type=code&client_id={{ client_id }}"
		baseURL += "&device_id={{ device_id }}&device_name=yapusher-cli&force_confirm=yes"

		baseURL = strings.Replace(baseURL, "{{ client_id }}", YANDEX_APPID, 1)
		baseURL = strings.Replace(baseURL, "{{ device_id }}", c.Config.DeviceID, 1)

		dlog.Info().Msg("Please open in your browser this URL and authorize the app. After getting the code restart the app with -authCode param (see -h for details).")
		dlog.Info().Msgf("Auth URL: %s", baseURL)

		os.Exit(0)
	}
}

// checkAuth detects if we have authorized already
func checkAuth() bool {
	if c.Config.Token.AccessToken != "" {
		return true
	}

	return false
}
