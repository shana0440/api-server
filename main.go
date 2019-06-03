package main

import (
	"github.com/dannypsnl/rocket"
	log "github.com/sirupsen/logrus"

	"github.com/u-job/api-server/config"
	"github.com/u-job/api-server/handlers"
)

func main() {
	configuration, err := config.LoadAndGetConfig()

	if err != nil {
		log.Fatalln("failed to load .env")
	}

	log.SetLevel(configuration.App.LogLevel)

	rocket.Ignite(configuration.App.Port).
		Mount(handlers.HelloWorldHandler).
		Launch()
}
