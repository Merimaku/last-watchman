package main

import (
	"github.com/Merimaku/last-watchman/pkg/config"
	"log"

	"github.com/Merimaku/last-watchman/pkg/watchman"
)

func main() {
	log.Println("Running")
	serverConfig, err := config.ReadConfigFromFile("conf-dev.toml")
	if err != nil {
		log.Fatalln(err)
	}
	watchman, err := watchman.AppBuilder(serverConfig)
	if err != nil {
		log.Fatalln(err)
	}
	watchman.Serve()
}