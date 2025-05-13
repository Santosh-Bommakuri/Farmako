package main

import (
	"Farmako/cache"
	"Farmako/config"
	"Farmako/routes"
)

func main() {

	config.ConnectDB()
	cache.ConnectRedis()

	app := routes.Setup()

	app.Run(":9091")
}
