package main

import (
	_ "backend/docs"
	"backend/internal"
	"backend/internal/server"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"os"
)

//	@title			MSMC API
//	@version		0.1
//	@description	This is MSMC API.

// @host		localhost:3000
// @BasePath	/
func main() {
	err := internal.Prepare()
	handleErr(err, "Failed to prepare the MSMC server")

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Route("/server", server.Router)

	err = app.Listen(":3000")
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
