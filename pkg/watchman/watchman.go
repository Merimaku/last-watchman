package watchman

import (
	"time"
	"fmt"

	"github.com/Merimaku/last-watchman/pkg/config"
)

type Watchman struct {
	Config *config.Watchman
}

func (w *Watchman) Server() {
	watchTicker := time.NewTicker(10 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
					return
			case t := <-watchTicker.C:
					fmt.Println("Tick at", t)
			}
		}
	}()
}