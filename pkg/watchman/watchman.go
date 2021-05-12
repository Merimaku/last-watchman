package watchman

import (
	"context"
	"time"
	"log"
	"os"
	"os/signal"
	// "net/http"

	"github.com/Merimaku/last-watchman/pkg/config"
)

type Watchman struct {
	Config 	*config.Watchman
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
			os.Exit(2)
	}()
	if err := run(ctx, watchTicker); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, ticker *time.Ticker) error {
	for {
		select {
		case <-ctx.Done():
			time.Sleep(5 * time.Second)
			log.Println("Terminating")
			return nil
		case t := <-(*ticker).C:
			log.Println("Tick at", t)
		}
	}
}