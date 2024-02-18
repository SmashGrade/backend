package main

import (
	"github.com/SmashGrade/backend/app/api"
	c "github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/db"
	_ "github.com/SmashGrade/backend/app/docs"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "gorm.io/gorm"
)

//	@title						Smashgrade Backend API
//	@version					1.0
//	@description				Backend API for Smashgrade, a web application for tracking your student grades.
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				HFTM Grenchen
//	@contact.url				https://www.hftm.ch
//	@license.name				Closed
//	@host						api.smashgrade.ch
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
func main() {

	// Load configuration
	config := c.NewAPIConfig()

	server := echo.New()
	// Assign the custom error handler to the server
	server.HTTPErrorHandler = e.HandleEchoError

	// Add swagger documentation route
	server.GET("/docs*", echoSwagger.WrapHandler)

	// Initialize the database provider
	provider := db.NewProvider(config)

	// Initialize the router
	router := api.NewRouter(server, provider)
	// Register all v1 routes
	router.RegisterV1()

	// Start the server
	// Any returned error is fatal
	server.Logger.Fatal(server.Start(config.ServerAddress()))

}
