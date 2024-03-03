package cmd

import (
	"github.com/SmashGrade/backend/app/api"
	c "github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the SmashGrade server",
	Long:  `Starts the SmashGrade server. The server is the main component of the SmashGrade backend. It provides a RESTful API for the frontend and handles all the business logic.`,
	Run:   StartServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("config", "c", "config.yaml", "Path to the configuration file")
}

func StartServer(cmd *cobra.Command, args []string) {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		configPath = "config.yaml"
	}

	// Load configuration
	config := c.FromFile(configPath)
	// Show the branding banner
	config.ShowBrand()

	// Create a new echo server
	server := echo.New()

	// Remove echo banner
	server.HideBanner = true
	server.HidePort = true

	// Assign the custom error handler to the server
	server.HTTPErrorHandler = e.HandleEchoError

	// Add swagger documentation route
	server.GET("/docs/*", echoSwagger.WrapHandler)

	// Enable Middleware
	server.Use(middleware.RequestLoggerWithConfig(config.GetEchoLoggerConfig()))
	server.Use(middleware.CORSWithConfig(config.GetEchoCORSConfig()))
	server.Use(middleware.BodyLimit(config.MaxBodySize))
	server.Use(middleware.RateLimiterWithConfig(config.GetRateLimitConfig()))

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
