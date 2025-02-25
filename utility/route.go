package utility

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(base *fiber.App) {
	// index route
	base.Get("/", Index)

	login := base.Group("/login")

	// login route
	login.Get("/", LoginPage)
	login.Post("/auth", LoginAuth)

	// dashboard admin route
	base.Get("/dashboard", Dashboard)
}
