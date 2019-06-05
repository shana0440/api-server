package main

import (
	"github.com/dannypsnl/rocket"
	log "github.com/sirupsen/logrus"

	"github.com/u-job/api-server/config"
	"github.com/u-job/api-server/db"
	"github.com/u-job/api-server/handlers"
)

func main() {
	configuration, err := config.LoadAndGetConfig()
	if err != nil {
		log.Fatalln("failed to load .env", err)
	}

	err = db.CreateConn(configuration.DB)
	if err != nil {
		log.Fatalln("failed to connect with database", err)
	}
	defer db.Close()

	log.SetLevel(configuration.App.LogLevel)

	rocket.Ignite(configuration.App.Port).
		Mount(handlers.CreateUser).
		Mount(handlers.GetUser).
		Launch()
}
