// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package yandexv1

type TokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
