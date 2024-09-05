package internal

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"strconv"
)

// requires at least one property, returns all of them
type userInfoDto struct {
	Id       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

// UserInfoController swagger:
//
//	@Summary Get user info
//	@Description Get user info
//	@Accept json
//	@Produce json
//	@Param id body int false "ID"
//	@Param email body string false "Email"
//	@Param username body string false "Username"
//	@Success 200 {object} userInfoDto
//	@Success 404 {object} object{error=string}
//	@Router /get-user-info [post]
func UserInfoController(ctx *fiber.Ctx) error {
	var dto userInfoDto
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// At least one userInfoDto field must be filled:
	if dto.Id == 0 && dto.Email == "" && dto.Username == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "At least one field must be filled",
		})
	}

	db := ctx.Locals("db").(*pgx.Conn)

	res, err := userInfo(db, dto)
	if errors.Is(err, pgx.ErrNoRows) {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Failed to get user info",
		})
	}

	return ctx.JSON(res)
}

func userInfo(db *pgx.Conn, dto userInfoDto) (userInfoDto, error) {
	var dbRes userInfoDto
	ctx := context.Background()

	err := db.QueryRow(ctx, `
		SELECT id, email, username FROM "users"
		WHERE id = $1 OR email = $2 OR username = $3
		LIMIT 1
	`, strconv.Itoa(dto.Id), dto.Email, dto.Username).Scan(&dbRes.Id, &dbRes.Email, &dbRes.Username)

	return dbRes, err
}
