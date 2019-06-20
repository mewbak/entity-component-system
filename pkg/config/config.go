// Package config contains methods and types for reading configuration
// data for the game.
package config

import "gopkg.in/ini.v1"

const (
	configName = "config.ini"
)

type (
	// The Config type represents all game configuration options
	Config struct {
		Width      int  // The screen width
		Height     int  // The screen height
		FullScreen bool // If true, runs in full screen mode
	}
)

// Load attempts to load the .ini file and parse its configuration
// options.
func Load() (*Config, error) {
	cfg, err := ini.Load(configName)

	if err != nil {
		return nil, err
	}

	out := &Config{}

	if err := parseGraphics(cfg, out); err != nil {
		return nil, err
	}

	return out, nil
}

func parseGraphics(in *ini.File, out *Config) error {
	gfx := in.Section("graphics")

	width, err := gfx.Key("width").Int()

	if err != nil {
		return err
	}

	height, err := gfx.Key("height").Int()

	if err != nil {
		return err
	}

	fs, err := gfx.Key("full_screen").Bool()

	if err != nil {
		return err
	}

	out.Width = width
	out.Height = height
	out.FullScreen = fs

	return nil
}
