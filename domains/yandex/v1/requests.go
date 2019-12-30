// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package yandexv1

import (
	"encoding/json"
	"github.com/schollz/progressbar/v2"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	form.Set("device_name", DefaultDeviceName)
	form.Set("client_id", YandexAppID)
	form.Set("client_secret", YandexAppPw)

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

		dlog.Error().Str("error", errorData.Error).Str("description", errorData.ErrorDescription).Msg("Got error from Yandex, not authorized. Please retry authorization")
		authorize()
	}

	os.Exit(0)
}

func uploadFile(uploadPath string, filePath string, overwriteFile bool) {
	uploadRequestURL := "https://cloud-api.yandex.net/v1/disk/resources/upload"

	// Checking file existence before requesting
	normalizedFilePath, _ := filepath.Abs(filePath)

	fileInfo, err := os.Stat(normalizedFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			dlog.Fatal().Err(err).Msg("File for uploading not found")
		} else {
			dlog.Fatal().Err(err).Msg("Failed to stat uploading file")
		}
	}

	if fileInfo.Size() > (MaxUploadSize - 1) {
		dlog.Fatal().Msg("Requested file is too big")
	}

	if !fileInfo.Mode().IsRegular() {
		dlog.Fatal().Msg("Only regular files uploading is supported right now")
	}

	client := http.Client{}
	uploadInfo := UploadInfo{}

	// The first request will get from Yandex upload URL
	req, _ := http.NewRequest("GET", uploadRequestURL, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "OAuth "+c.Config.Token.AccessToken)

	query := url.Values{}
	query.Add("path", "disk:/"+uploadPath+"/"+fileInfo.Name())
	query.Add("overwrite", strconv.FormatBool(overwriteFile))

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		dlog.Fatal().Err(err).Msg("Failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&uploadInfo)
		if err != nil {
			dlog.Error().Err(err).Msg("Failed to decode response")
		}
	} else {
		errorData := UploadError{}
		err = json.NewDecoder(resp.Body).Decode(&errorData)
		if err != nil {
			dlog.Error().Err(err).Msg("Failed to decode response")
		}

		dlog.Info().Str("error", errorData.Error).Str("description", errorData.Description).Msg("Failed to upload file")
		os.Exit(1)
	}

	file, _ := os.Open(normalizedFilePath)
	bar := progressbar.NewOptions(
		int(fileInfo.Size()),
		progressbar.OptionSetBytes(int(fileInfo.Size())),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionClearOnFinish(),
	)

	progressReader := &progressReader{
		r:           file,
		progressbar: bar,
	}

	if uploadInfo.URL == "" {
		dlog.Fatal().Msg("Got empty upload URL. Report a bug at https://source.hodakov.me/fat0troll/yapusher/issues because this situation is impossible.")
	}

	uploadReq, _ := http.NewRequest("PUT", uploadInfo.URL, progressReader)

	uploadResp, err := client.Do(uploadReq)
	if err != nil {
		dlog.Fatal().Err(err).Msg("Failed to send upload request")
	}
	defer uploadResp.Body.Close()

	switch uploadResp.StatusCode {
	case http.StatusCreated:
		dlog.Info().Msg("File uploaded successfully")
	case http.StatusAccepted:
		dlog.Info().Msg("File uploaded successfully, but it will take time for Yandex.Disk to handle it internally. Be patient and don't try to upload single file many times")
	case http.StatusRequestEntityTooLarge:
		dlog.Fatal().Msg("File upload is too large. Report a bug at https://source.hodakov.me/fat0troll/yapusher/issues because this situation should be handled before upload attempt.")
	case http.StatusInsufficientStorage:
		dlog.Fatal().Msg("There is no space left on your Yandex.Disk.")
	default:
		dlog.Fatal().Msg("Failed to upload file (error on Yandex's side). Try again later.")
	}

	os.Exit(0)
}
