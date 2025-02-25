package main

import (
	"github.com/deafzero/LoginTest/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	base := fiber.New(fiber.Config{
		Views: html.New("views", ".html"),
	})

	// register route and hanlder
	utility.RegisterRoute(base)

	// serve server
	base.Listen(":8080")
}
