package bootstrap

import (
	"go-fiber-minimal/app/middleware"
	"go-fiber-minimal/lib"
	"go-fiber-minimal/route"

	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	lib.Env.Init()

	lib.LogFile.Init()

	lib.DB.Init()

	route.Console.Init()

	lib.Validator.Init()

	app := lib.Fiber.Init()
	app.Use(middleware.Cors.Init())

	route.Api(app)
	route.Web(app)

	return app
}
