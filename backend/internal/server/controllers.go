package server

import (
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/", getServersController)
	router.Post("/", createServerController)
}

// createServerController
//
//	@Summary		Create a new server
//	@Tags			server
//	@ID				create-server
//	@Accept			json
//	@Produce		json
//	@Param			server	body createServerDto	true	"Server data"
//	@Success		200	{string}	string			"Server created successfully"
//	@Failure		400	{string}	string			"Bad request"
//	@Failure		500	{string}	string			"Internal server error"
//	@Router			/server [post]
func createServerController(ctx *fiber.Ctx) error {
	var dto createServerDto
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	apiErr := createServer(dto)
	if apiErr.IsNotNil() {
		return apiErr.Send(ctx)
	}

	return ctx.SendString("Server created successfully")
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
