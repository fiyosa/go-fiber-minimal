package bootstrap

import (
	"go-fiber-ddd/config"
	"go-fiber-ddd/lib"
	"go-fiber-ddd/route"

	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	lib.Env.Init()

	lib.LogFile.Init()

	lib.DB.Init()

	route.Console.Init()

	lib.Validator.Init()

	app := lib.Fiber.Init()
	app.Use(config.Cors())

	lib.LogConsole.Info("Port: " + config.Env.APP_PORT)

	return app
}
