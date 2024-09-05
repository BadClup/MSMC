package internal

import (
	"backend/shared"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"math"
	"strings"
)

type JwtPayload struct {
	Exp    int64 `json:"ExpiresAt"`
	UserID int   `json:"user_id"`
}

type LoginRemoteDto struct {
	UserId int `json:"user_id"`
}

// LoginRemoteController swagger:
//
//	@Summary Login handler for auth-service only
//	@Description This is a handler for auth-service only. It returns a token, which should be forwarded to the client.
//	@Tags auth
//	@Accept json
//	@Produce json
//	@Param body body loginRemoteDto true "User ID"
//	@Success 200 {object} object{token=string}
//	@Failure 400 {object} object{error=string}
//	@Failure 401 {object} object{error=string}
//	@Router /login-remote [post]
func LoginRemoteController(ctx *fiber.Ctx) error {
	var dto LoginRemoteDto

	hostnameWithoutPort := ctx.Hostname()
	hostnameWithoutPort = strings.Split(hostnameWithoutPort, ":")[0]

	// Validate if request comes from authorized server:
	if hostnameWithoutPort != shared.AuthServiceHostname {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fmt.Sprintf("Hostname %s is not authorized", hostnameWithoutPort),
		})
	}

	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request " + string(ctx.BodyRaw()),
		})
	}

	token, err := signJwt(JwtPayload{
		Exp:    math.MaxInt32,
		UserID: dto.UserId,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

type GetJwtPayloadDto struct {
	Token string `json:"token"`
}

// GetJwtPayloadController swagger:
//
//	@Summary Get JWT payload
//	@Description This is a handler for auth-service only.
//	@Tags auth
//	@Accept json
//	@Produce json
//	@Param body body GetJwtPayloadDto true "Token"
//	@Success 200 {object} object{payload=JwtPayload}
//	@Failure 400 {object} object{error=string}
//	@Router /get-jwt-payload [post]
func GetJwtPayloadController(ctx *fiber.Ctx) error {
	var dto GetJwtPayloadDto
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	payload, err := decodeJwt(dto.Token)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	return ctx.JSON(fiber.Map{
		"payload": payload,
	})
}

func signJwt(payload JwtPayload) (string, error) {
	if payload.Exp == 0 {
		payload.Exp = math.MaxInt32
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": payload.Exp,
		"user_id":   payload.UserID,
	})

	return token.SignedString([]byte(shared.JwtSecret))
}

func decodeJwt(tokenString string) (JwtPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(shared.JwtSecret), nil
	})
	if err != nil {
		return JwtPayload{}, err
	}

	if !token.Valid {
		return JwtPayload{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return JwtPayload{}, errors.New("invalid token claims")
	}

	return JwtPayload{
		Exp:    int64(claims["ExpiresAt"].(float64)),
		UserID: int(claims["user_id"].(float64)),
	}, nil
}
