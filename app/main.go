package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// TODO: https://github.com/swaggo/echo-swagger

// @title Smashgrade Backend API
// @version 1.0
// @description Backend API for Smashgrade, a web application for tracking your student grades.
// @termsOfService http://swagger.io/terms/
// @contact.name HFTM Grenchen
// @contact.url https://www.hftm.ch
// @license.name Closed
// @host api.smashgrade.ch
// @BasePath /v1
func main() {

	// Load configuration
	config := NewAPIConfig()

	server := echo.New()
	// Assign the custom error handler to the server
	server.HTTPErrorHandler = HandleEchoError

	// Add swagger documentation route
	server.GET("/docs/*", echoSwagger.WrapHandler)

	// Start the server
	// Any returned error is fatal
	server.Logger.Fatal(server.Start(config.ServerAddress()))

}
