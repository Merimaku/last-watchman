package config

import (

	"github.com/BurntSushi/toml"
)

type LastFM struct {
	APIKey				string 	`toml:"api_key"`
	ClientSecret	string 	`toml:"secret"`
}

type Watchman struct {
	LastFM	*LastFM	`toml:"lastfm"`
}

func ReadConfigFromFile(filePath string) (*Watchman, error) {
	watchman := &Watchman{}
	if _, err := toml.DecodeFile(filePath, watchman); err != nil {
		return nil, err
	}
	return watchman, nil
}