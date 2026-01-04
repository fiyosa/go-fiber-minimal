package main

import (
	"go-fiber-minimal/bootstrap"
	"go-fiber-minimal/config"
)

func main() {
	f := bootstrap.Init()

	f.Listen(":" + config.Env.APP_PORT)
}
