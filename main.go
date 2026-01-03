package main

import (
	"go-fiber-ddd/bootstrap"
	"go-fiber-ddd/config"
	"go-fiber-ddd/lib"
)

func xmain() {
	lib.Env.Init()

	lib.LogFile.Init()

	lib.LogFile.Info("APP_PORT", config.Env.APP_PORT)
}

func main() {
	f := bootstrap.Init()

	f.Listen(":" + config.Env.APP_PORT)
}
