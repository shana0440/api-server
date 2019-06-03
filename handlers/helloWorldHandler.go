package handlers

import (
	"github.com/dannypsnl/rocket"
	log "github.com/sirupsen/logrus"
)

// HelloWorldHandler a example for how to use rocket
var HelloWorldHandler = rocket.Get("/", func() string {
	log.Debug("hello")
	return "hello world"
})
