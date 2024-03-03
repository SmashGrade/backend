package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the root command for the application
var rootCmd = &cobra.Command{
	Use:   "smashgrade",
	Short: "SmashGrade is a web application for tracking your student grades.",
	Long:  `SmashGrade is a web application for tracking your student grades. The backend provides a RESTful API for the frontend and handles all the business logic.`,
	// Without any subcommands, the root command will start the server
	Run: StartServer,
}

// Execute is the entry point for the root command
func Execute() {
	// Run the root command
	err := rootCmd.Execute()
	// Handle any errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Initialize the root command
func Init() {
	rootCmd.Flags().StringP("config", "c", "config.yaml", "Path to the configuration file")
}
