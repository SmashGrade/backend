package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smashgrade",
	Short: "SmashGrade is a web application for tracking your student grades.",
	Long:  `SmashGrade is a web application for tracking your student grades. The backend provides a RESTful API for the frontend and handles all the business logic.`,
	Run:   StartServer,
}

// Execute is the entry point for the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
