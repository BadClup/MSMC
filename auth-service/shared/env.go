package shared

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
)

func warn(msg string) {
	log.Warnf("Production only warning: %s", msg)
}

var DatabaseUrl = func() string {
	pgUrl := os.Getenv("DATABASE_URL")
	if pgUrl != "" {
		return pgUrl
	}
	warn("Using default database URL")
	return "postgres://postgres:postgres@localhost:5432/msmc-auth-service"
}()

var JwtSecret = func() string {
	secret := os.Getenv("JWT_SECRET")
	if secret != "" {
		return secret
	}
	warn("Using default JWT secret")
	return "9yoh7g8kyuv7t6c5r6x5e4w3q2a1z"
}()
