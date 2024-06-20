package main

import (
	"myapp/internal/core"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

// main is the entry point of the application. It initializes the application
// by creating an instance of the core app and starting it.
// In most cases there should be no need to modify this function.
func main() {
	app := core.CreateApp()

	if err := app.Start(); err != nil {
		app.Logger().Error(err.Error())
		os.Exit(1)
	}
}
