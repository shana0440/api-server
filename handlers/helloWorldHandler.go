package handlers

import (
	"github.com/dannypsnl/rocket"
)

// HelloWorldHandler a example for how to use rocket
var HelloWorldHandler = rocket.Get("/", func () string {
	return "hello world"
})