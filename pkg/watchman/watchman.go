package watchman

import (
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

func (w *Watchman) Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	watchTicker := time.NewTicker(1 * time.Second)

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
	if err := runForever(ctx, watchTicker); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func runOnce() error {
	return nil
}

func runForever(ctx context.Context, ticker *time.Ticker) error {
	for {
		select {
		case <-ctx.Done():
			time.Sleep(5 * time.Second)
			log.Println("Terminating")
			return nil
		case <-(*ticker).C:
			if err := runOnce(); err != nil {
				log.Fatalln(err)
			}
		}
	}
}