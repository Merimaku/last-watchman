package watcher

import (
	"github.com/Merimaku/last-watchman/pkg/modules/lastfm"
	"github.com/Merimaku/last-watchman/pkg/config"
)

type Service struct {
	config	*config.Service
	lastfmService  *lastfm.Service
}

// NewService returns a Service that contains the config struct with the service level configuration
// and lastfm module's Service struct
func NewService(
	config	*config.Service,
	lastfmService  *lastfm.Service,
) (*Service, error) {
	service := &Service{
		config:	config,
		lastfmService: lastfmService,
	}

	return service, nil
}

func (w *Service) Run() error {

	return nil
}