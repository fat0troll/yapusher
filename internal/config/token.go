// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package config

type Token struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpireTime   int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
