package internal

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"math"
	"msmc/auth-server/shared"
)

// RegisterController swagger:
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
	var dto registerDto
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db := shared.GetDatabaseConn(ctx)

	token, err := registerUser(db, dto)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

type registerDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type jwtPayload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
}

// registerUser return jwt token
func registerUser(db *pgx.Conn, dto registerDto) (string, error) {
	// TODO: implement user registration

	var userId int

	ctx := context.Background()
	err := db.QueryRow(ctx, `
		INSERT INTO "users" (email, username, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, dto.Email, dto.Username, dto.Password).Scan(&userId)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      math.MaxInt64,
		"email":    dto.Email,
		"username": dto.Username,
		"user_id":  userId,
	})

	return token.SignedString([]byte(shared.JwtSecret))
}
