package server

import (
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/", getServersController)
}

// getServersController
//
//	@Summary		Get all servers
//	@Tags			server
//	@ID				get-all-servers
//	@Produce		json
//	@Success		200	{array}	shared.ServerInstanceStatus
//	@Failure		500	{string}	error
//	@Router			/server [get]
func getServersController(ctx *fiber.Ctx) error {
	servers, err := getAllServers()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.JSON(servers)
}
