package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-CSRF-Token, Cache-Control, X-Requested-With",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH, OPTIONS",
		MaxAge:           86400, // 24 hours
	})
}
