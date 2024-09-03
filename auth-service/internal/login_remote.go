package internal

import (
	"context"
	"github.com/carlmjohnson/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strings"
)

type loginRemoteDto struct {
	Token     string `json:"token"`
	RemoteUrl string `json:"url"`
}

// LoginRemoteController swagger:
//
//	@Summary Login a user remotely
//	@Description Login a user to remote server
//	@Accept json
//	@Produce json
//	@Param token body string true "Token"
//	@Param url body string true "URL"
//	@Success 200 {object} tokenResponse
//	@Router /login-remote [post]
func LoginRemoteController(ctx *fiber.Ctx) error {
	var payload loginRemoteDto
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if payload.Token == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Token is required",
		})
	}

	if payload.RemoteUrl == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "URL is required",
		})
	}

	tokenPayload, err := decodeJwt(payload.Token)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	token, err := loginRemote(payload.RemoteUrl, tokenPayload.UserID)
	if err != nil {
		log.Debugf("Failed to login remote: %v", err)
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Failed to login",
		})
	}

	return ctx.JSON(token)
}

func loginRemote(url string, userId int) (tokenResponse, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	ctx := context.Background()
	var res tokenResponse

	err := requests.
		URL(url).
		Path("/login-remote").
		Post().
		BodyJSON(map[string]int{
			"user_id": userId,
		}).
		ToJSON(res).
		Fetch(ctx)

	return res, err
}
