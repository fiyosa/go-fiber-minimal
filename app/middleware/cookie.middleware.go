package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var Cookie cookieManager

type cookieManager struct{}

func (cookieManager) Set(c *fiber.Ctx, name string, value string) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HTTPOnly: true, // true: can't access from js
		Secure:   true, // true: https & false: http
		SameSite: fiber.CookieSameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})
}

func (cookieManager) Get(c *fiber.Ctx, name string) string {
	return c.Cookies(name)
}

func (cookieManager) Remove(c *fiber.Ctx, name string) {
	c.ClearCookie(name)
}
