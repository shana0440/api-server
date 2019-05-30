package main

import (
	"github.com/dannypsnl/rocket"
	"github.com/u-job/api-server/handlers"
)

func main() {
	rocket.Ignite(":3000").
		Mount(handlers.HelloWorldHandler).
		Launch()
}
