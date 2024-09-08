package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "msmc/auth-service/docs"
	"msmc/auth-service/internal"
	"msmc/auth-service/shared"
	"os"
	"time"
)

//	@title			MSMC auth API
//	@version		0.1
//	@description	This is the auth server for MSMC

// @host		localhost:3001
// @BasePath	/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file, using default values")
	}

	db, err := dbConnect(10, 1000) // TODO: move these params to env
	handleErr(err, "Failed to connect to database")

	app := internal.GetApp(db)

	err = app.Listen(":3001")
	handleErr(err, "Failed to start server on port 3000")
}

// dbConnect - set maxAttempts to -1 for infinite attempts
func dbConnect(maxAttempts int, delayMs int) (*pgx.Conn, error) {
	var db *pgx.Conn
	var err error

	for i := 0; i < maxAttempts; i++ {
		ctx := context.Background()
		db, err = pgx.Connect(ctx, shared.DatabaseUrl)

		if err == nil {
			return db, err
		}
		fmt.Printf("Failed to connect to database, retrying in %d ms for the %d time\n", delayMs, i+1)
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}

	return nil, err
}

// only for top-level usage
func handleErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
