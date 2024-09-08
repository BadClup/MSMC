package test

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"github.com/jackc/pgx/v5"
	"msmc/auth-service/shared"
	"testing"
)

func GetDb(t *testing.T) *pgx.Conn {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, shared.DatabaseUrl)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

type UserDto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var SeedUsers = []UserDto{
	{
		Email:    "test@example.com",
		Username: "test",
		Password: "test",
	},
}

func SeedDb(t *testing.T) {
	ctx := context.Background()

	for _, user := range SeedUsers {
		passwordHashed := sha512.New()
		passwordHashed.Write([]byte(user.Password))

		println("passwordHashed.Sum(nil): ", hex.EncodeToString(passwordHashed.Sum(nil)))

		_, err := GetDb(t).Exec(ctx, `
			INSERT INTO "users" (email, username, password)
			VALUES ($1, $2, $3)
		`, user.Email, user.Username, hex.EncodeToString(passwordHashed.Sum(nil)))
		if err != nil {
			t.Fatalf("Failed to seed database: %v", err)
		}
	}
}

func ClearDb(t *testing.T) {
	ctx := context.Background()

	_, err := GetDb(t).Exec(ctx, `
		TRUNCATE TABLE "users"
	`)
	if err != nil {
		t.Fatalf("Failed to clear database: %v", err)
	}
}
