package watchman

import (
	"github.com/Merimaku/last-watchman/pkg/modules/lastfm"
	"github.com/Merimaku/last-watchman/pkg/modules/watcher"
	"github.com/Merimaku/last-watchman/pkg/config"
)

// Return a struct that contains the configuration object,
// watcherService object that run the watcher service
type Modules struct {
	config	*config.Watchman

	watcherService *watcher.Service
}

func initialiseModules(config *config.Watchman) (Modules, error) {
	lastfmService := lastfm.NewService(config.LastFM)
	watcherService, err := watcher.NewService(config.Service, lastfmService)
	if err != nil {
		return Modules{}, err
	}
	modules := Modules{
		config:	config,
		watcherService: watcherService,
	}
	return modules, nil
}