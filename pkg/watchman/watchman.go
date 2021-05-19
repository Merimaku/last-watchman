package watchman

import (
	"syscall"
	"context"
	"time"
	"log"
	"os"
	"os/signal"
	"errors"

	"github.com/Merimaku/last-watchman/pkg/config"
)

type Watchman struct {
	Config 	*config.Watchman
	Modules *Modules
}

// Return the Watchman object that contains the config and initialise all necessary services
func AppBuilder(config *config.Watchman) (*Watchman, error) {
	if config == nil {
		return nil, errors.New("Config is not set")
	}
	modules, err := initialiseModules(config)
	if err != nil {
		return nil, err
	}
	return &Watchman{
		Config: config,
		Modules: &modules,
	}, nil
}

// Start the Watchman app main service loop
func (w *Watchman) Serve() error {
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	watchTicker := time.NewTicker(w.Config.Service.UpdateInterval.Duration)

	go func() {
			select {
			case <-signalChan:
				log.Println("System Interrupted")
				cancel()
			case <-ctx.Done():
			}
			<-signalChan
			log.Println("Hard Exit")
			os.Exit(1)
	}()
	if err := w.runForever(ctx, watchTicker); err != nil {
		return err
	}
	return nil
}

// runOnce starts every services in the Modules object
func (w *Watchman) runOnce() error {
	err := w.Modules.watcherService.Run()
	if err != nil {
		return err
	}
	return nil
}

// runForever wrap runOnce around a context switch loop
func (w *Watchman) runForever(ctx context.Context, ticker *time.Ticker) error {
	for {
		select {
		case <-ctx.Done():
			time.Sleep(5 * time.Second)
			log.Println("Terminating")
			return nil
		case <-(*ticker).C:
			if err := w.runOnce(); err != nil {
				return err
			}
		}
	}
}