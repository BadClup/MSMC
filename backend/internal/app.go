package internal

import (
	"backend/internal/mcserver"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"os"
)

// GetApp This function is only used by main.go and tests
// Because of that it can exit the program on error
func GetApp() *fiber.App {
	err := Prepare()
	handleErr(err, "Failed to prepare the MSMC server")

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Route("/server", mcserver.Router)
	app.Post("/login-remote", LoginRemoteController)
	app.Post("/get-jwt-payload", GetJwtPayloadController)

	return app
}

// only for top-level usage
func handleErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
