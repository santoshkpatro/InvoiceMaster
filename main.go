package main

import (
	"InvoiceMaster/cmd"
	"InvoiceMaster/config"
	"InvoiceMaster/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Check if there are command-line arguments.
	if len(os.Args) > 1 {
		cmd.Execute() // Run the CLI command
		return
	}

	// Connect to Database
	config.ConnectDB()

	// If no CLI commands, start the HTTP server.
	startServer()
}

func startServer() {
	e := echo.New()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
