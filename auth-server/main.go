package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "msmc/auth-server/docs"
	"msmc/auth-server/internal"
	"msmc/auth-server/shared"
	"os"
)

//	@title			MSMC auth server API
//	@version		0.1
//	@description	This is the auth server for MSMC

// @host		localhost:3001
// @BasePath	/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file, using default values")
	}

	ctx := context.Background()
	db, err := pgx.Connect(ctx, shared.DatabaseUrl)
	handleErr(err, "Failed to connect to database")

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/register", internal.RegisterController)

	err = app.Listen(":3001")
	handleErr(err, "Failed to start server on port 3000")
}

// only for top-level usage
func handleErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
