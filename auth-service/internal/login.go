package internal

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"math"
)

type LoginDto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginController swagger:
//
//	@Summary Login a user
//	@Description Login a user
//	@Accept json
//	@Produce json
//	@Param email body string false "Email"
//	@Param username body string false "Username"
//	@Param password body string true "Password"
//	@Success 200 {object} AuthResponse
//	@Router /login [post]
func LoginController(ctx *fiber.Ctx) error {
	var payload LoginDto
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if payload.Email == "" && payload.Username == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Email or username is required",
		})
	}

	if payload.Password == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Password is required",
		})
	}

	db := ctx.Locals("db").(*pgx.Conn)

	response, err := login(db, payload)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Message,
		})
	}

	return ctx.JSON(response)
}

func login(db *pgx.Conn, payload LoginDto) (AuthResponse, *fiber.Error) {
	var userId int
	var email, username, password string

	ctx := context.Background()
	err := db.QueryRow(ctx, `
		SELECT id, email, username, password
		FROM users
		WHERE email = $1 OR username = $2
	`, payload.Email, payload.Username).Scan(&userId, &email, &username, &password)

	if err != nil {
		log.Debugf("Failed to select user while loggin in: %v", err)
		return AuthResponse{}, fiber.NewError(400, "invalid email or username")
	}

	hashedInputPassword := sha512.New()
	hashedInputPassword.Write([]byte(payload.Password))

	if password != hex.EncodeToString(hashedInputPassword.Sum(nil)) {
		return AuthResponse{}, fiber.NewError(400, "invalid password")
	}

	token, err := signJwt(jwtPayload{
		Exp:      math.MaxInt64,
		Email:    email,
		Username: username,
		UserID:   userId,
	})
	if err != nil {
		return AuthResponse{}, fiber.NewError(500, "Failed to sign jwt")
	}

	return AuthResponse{
		Token: token,
		User: UserPublicData{
			Email:    email,
			Username: username,
			Id:       userId,
		},
	}, nil
}
