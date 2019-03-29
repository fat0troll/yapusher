// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package config

type Config struct {
	DeviceID string `json:"device_id"`
	Token    Token  `json:"token,omitempty"`
}
