package internal

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"math"
)

type jwtPayload struct {
	Exp    int64 `json:"exp"`
	UserID int   `json:"user_id"`
}

type loginRemoteDto struct {
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
	var dto loginRemoteDto

	// Validate if request comes from authorized server:
	// TODO: replace with hostname from .env
	if ctx.Hostname() != "localhost" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	token, err := signJwt(jwtPayload{
		Exp:    math.MaxInt64,
		UserID: dto.UserId,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func signJwt(payload jwtPayload) (string, error) {
	if payload.Exp == 0 {
		payload.Exp = math.MaxInt64
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     payload.Exp,
		"user_id": payload.UserID,
	})

	// TODO: replace with secret from .env
	return token.SignedString([]byte("TODO"))
}

func decodeJwt(tokenString string) (jwtPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// TODO: replace with secret from .env
		return []byte("TODO"), nil
	})
	if err != nil {
		return jwtPayload{}, err
	}

	if !token.Valid {
		return jwtPayload{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwtPayload{}, errors.New("invalid token claims")
	}

	return jwtPayload{
		Exp:    int64(claims["exp"].(float64)),
		UserID: int(claims["user_id"].(float64)),
	}, nil
}
