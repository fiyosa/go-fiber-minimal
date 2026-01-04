package route

import (
	"go-fiber-minimal/app/http/controller"
	"go-fiber-minimal/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	r := app.Group("/api")
	Auth := middleware.Auth

	r.Get("/auth/user", Auth(), controller.Auth.User)
	r.Post("/auth/login", controller.Auth.Login)
	r.Post("/auth/register", controller.Auth.Register)
	r.Post("/auth/logout", Auth(), controller.Auth.Logout)

	r.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "API not found"})
	})
}
