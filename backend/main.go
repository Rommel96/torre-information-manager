package main

import (
	"github.com/rommel96/torre-information-manager/backend/config"
	"github.com/rommel96/torre-information-manager/backend/server"
)

func main() {
	//server.Run()
	config.RunConfig()
	server.Run()
}
