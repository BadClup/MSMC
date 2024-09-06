package internal

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"strconv"
)

// UserInfoController swagger:
//
//	@Summary Get user info
//	@Description Get user info
//	@Accept json
//	@Produce json
//	@Param id body int false "ID"
//	@Param email body string false "Email"
//	@Param username body string false "Username"
//	@Success 200 {object} UserPublicData
//	@Success 404 {object} object{error=string}
//	@Router /get-user-data [post]
func UserInfoController(ctx *fiber.Ctx) error {
	var dto UserPublicData
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// At least one UserPublicData field must be filled:
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

func userInfo(db *pgx.Conn, dto UserPublicData) (UserPublicData, error) {
	var dbRes UserPublicData
	ctx := context.Background()

	err := db.QueryRow(ctx, `
		SELECT id, email, username FROM "users"
		WHERE id = $1 OR email = $2 OR username = $3
		LIMIT 1
	`, strconv.Itoa(dto.Id), dto.Email, dto.Username).Scan(&dbRes.Id, &dbRes.Email, &dbRes.Username)

	return dbRes, err
}
