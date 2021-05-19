package lastfm

import (
	"github.com/Merimaku/last-watchman/pkg/config"
)


type Service struct {
	config *config.LastFM
}

// NewService returns a Service that contains the LastFM struct with the LastFM API credential
func NewService(
	config *config.LastFM,
) (*Service) {
	return &Service{
		config:	config,
	}
}