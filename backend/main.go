package main

import (
	_ "backend/docs"
	"backend/internal"
	"fmt"
	"os"
)

//	@title			MSMC API
//	@version		0.1
//	@description	This is MSMC API.

// @host		localhost:3000
// @BasePath	/
func main() {
	app := internal.GetApp()

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Failed to start the server on port 3000")
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
