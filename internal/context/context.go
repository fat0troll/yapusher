// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package context

import (
	"encoding/json"
	"fmt"
	"github.com/kirsle/configdir"
	"github.com/rs/zerolog"
	"gitlab.com/pztrn/flagger"
	"gitlab.com/pztrn/go-uuid"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// getMemoryUsage returns memory usage for logger.
func (c *Context) getMemoryUsage(e *zerolog.Event, level zerolog.Level, message string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	e.Str("memalloc", fmt.Sprintf("%dMB", m.Alloc/1024/1024))
	e.Str("memsys", fmt.Sprintf("%dMB", m.Sys/1024/1024))
	e.Str("numgc", fmt.Sprintf("%d", m.NumGC))
}

// generateDefaultConfig generates new config on first run
func (c *Context) generateDefaultConfig() {
	newDeviceID, _ := uuid.NewV4()
	c.Config.DeviceID = newDeviceID.String()
}

// initFlagger initializes flags parser
func (c *Context) initFlagger() {
	c.Flagger = flagger.New("Yandex Disk Files Pusher", flagger.LoggerInterface(log.New(os.Stdout, "", log.Lshortfile)))
	c.Flagger.Initialize()
}

// Init is an initialization function for core context
// Without these parts of the application we can't start at all
func (c *Context) Init() {
	c.initFlagger()

	c.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	c.Logger = c.Logger.Hook(zerolog.HookFunc(c.getMemoryUsage))

	c.Logger.Info().Str("version", VERSION).Msg("yapusher is starting")

	dlog = c.Logger.With().Str("domain", "context").Logger()
}

func (c *Context) InitConfig() {
	configPath := configdir.LocalConfig("yapusher")
	err := configdir.MakePath(configPath)
	if err != nil {
		dlog.Fatal().Err(err).Str("config directory", configPath).Msg("Failed to obtain config path")
	}

	dlog.Debug().Str("config directory", configPath).Msg("Got config directory")

	configFile := filepath.Join(configPath, "settings.json")
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		// Generating new config on first run
		dlog.Debug().Msg("Generating new config")

		c.generateDefaultConfig()
		fh, err := os.Create(configFile)
		if err != nil {
			dlog.Fatal().Err(err).Msg("Failed to create config file")
		}
		defer fh.Close()

		encoder := json.NewEncoder(fh)
		err = encoder.Encode(c.Config)
		if err != nil {
			dlog.Fatal().Err(err).Msg("Failed to encode config")
		}
	} else {
		dlog.Debug().Msg("Using existing config")

		fh, err := os.Open(configFile)
		if err != nil {
			dlog.Fatal().Err(err).Msg("Failed to read config file")
		}

		defer fh.Close()

		decoder := json.NewDecoder(fh)
		decoder.Decode(c.Config)
		if err != nil {
			dlog.Fatal().Err(err).Msg("Failed to decode config")
		}
	}
}
