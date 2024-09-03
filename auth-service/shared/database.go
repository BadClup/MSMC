package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func GetDatabaseConn(ctx *fiber.Ctx) *pgx.Conn {
	return ctx.Locals("db").(*pgx.Conn)
}
