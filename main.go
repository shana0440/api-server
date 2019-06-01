package main

import (
	"log"

	"github.com/dannypsnl/rocket"

	"github.com/u-job/api-server/config"
	"github.com/u-job/api-server/handlers"
)

func main() {
	configuration, err := config.LoadAndGetConfig()

	if err != nil {
		log.Fatalln("failed to load .env")
	}

	rocket.Ignite(configuration.App.Port).
		Mount(handlers.HelloWorldHandler).
		Launch()
}
