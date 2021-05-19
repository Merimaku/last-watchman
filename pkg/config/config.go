package config

import (
	"github.com/BurntSushi/toml"
)

type Service struct {
	TimeOut duration 				`toml:"timeout"`
	UpdateInterval duration `toml:"update_interval"`
}

type LastFM struct {
	APIKey				string 	`toml:"api_key"`
	ClientSecret	string 	`toml:"secret"`
}

type Watchman struct {
	LastFM	*LastFM	 `toml:"lastfm"`
	Service *Service `toml:"service"`
}

func ReadConfigFromFile(filePath string) (*Watchman, error) {
	watchman := &Watchman{}
	if _, err := toml.DecodeFile(filePath, watchman); err != nil {
		return nil, err
	}
	return watchman, nil
}