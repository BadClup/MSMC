package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5"
)

func GetApp(db *pgx.Conn) *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/get-user-data", UserInfoController)
	app.Post("/register", RegisterController)
	app.Post("/login", LoginController)
	app.Post("/login-remote", LoginRemoteController)

	return app
}
