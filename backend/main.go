package main

import (
	_ "backend/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title MSMC API
// @version 0.1
// @description This is MSMC API.

// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
