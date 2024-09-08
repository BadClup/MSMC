package internal

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"math"
	"msmc/auth-service/shared"
)

type RegisterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// RegisterController swagger:
// TODO: implement RegisterDto validation (like password length)
//
//	@Summary Register a new user
//	@Description Register a new user
//	@Accept json
//	@Produce json
//	@Param email body string true "Email"
//	@Param password body string true "Password"
//	@Param username body string true "Username"
//	@Success 200 {object} object{token=string}
//	@Router /register [post]
func RegisterController(ctx *fiber.Ctx) error {
	var dto RegisterDto
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db := shared.GetDatabaseConn(ctx)

	response, err := registerUser(db, dto)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	return ctx.JSON(response)
}

// registerUser return jwt token
func registerUser(db *pgx.Conn, dto RegisterDto) (AuthResponse, error) {
	passwordHashed := sha512.New()
	passwordHashed.Write([]byte(dto.Password))

	var userId int

	ctx := context.Background()
	err := db.QueryRow(ctx, `
		INSERT INTO "users" (email, username, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, dto.Email, dto.Username, hex.EncodeToString(passwordHashed.Sum(nil))).Scan(&userId)
	if err != nil {
		return AuthResponse{}, err
	}

	token, err := signJwt(jwtPayload{
		Exp:      math.MaxInt64,
		Email:    dto.Email,
		Username: dto.Username,
		UserID:   userId,
	})
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{
		Token: token,
		User: UserPublicData{
			Email:    dto.Email,
			Username: dto.Username,
			Id:       userId,
		},
	}, nil
}
