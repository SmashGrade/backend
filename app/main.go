package main

import (
	c "github.com/SmashGrade/backend/app/config"
	_ "github.com/SmashGrade/backend/app/docs"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "gorm.io/gorm"
)

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
	config := c.NewAPIConfig()

	server := echo.New()
	// Assign the custom error handler to the server
	server.HTTPErrorHandler = e.HandleEchoError

	// Add swagger documentation route
	server.GET("/docs/*", echoSwagger.WrapHandler)

	// Start the server
	// Any returned error is fatal
	server.Logger.Fatal(server.Start(config.ServerAddress()))

}
