// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package yandexv1

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func sendCode(code int) {
	baseURL := "https://oauth.yandex.ru/token"

	client := http.Client{}
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", strconv.Itoa(code))
	form.Set("device_id", c.Config.DeviceID)
	form.Set("device_name", DEVICE_NAME)
	form.Set("client_id", YANDEX_APPID)
	form.Set("client_secret", YANDEX_APPPW)

	req, _ := http.NewRequest("POST", baseURL, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		dlog.Fatal().Err(err).Msg("Failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&c.Config.Token)
		if err != nil {
			dlog.Error().Err(err).Msg("Failed to decode response")
		}

		c.SaveConfig()

		dlog.Info().Msg("You've authorized successfully. Now you can use this app to upload files to your Yandex.Disk")
		dlog.Info().Msg("See -h for details")
	} else {
		errorData := TokenError{}
		err = json.NewDecoder(resp.Body).Decode(&errorData)
		if err != nil {
			dlog.Error().Err(err).Msg("Failed to decode response")
		}

		dlog.Error().Interface("response", errorData).Msg("Got error from Yandex, not authorized. Please retry authorization")
		authorize()
	}

	os.Exit(0)
}

func uploadFile() {}
