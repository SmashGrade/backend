package cmd

import (
	"fmt"
	"net/mail"
	"slices"
	"strings"

	c "github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"

	"github.com/spf13/cobra"
)

// Command to add a new user
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new user",
	Long:  `Add a new user to the database. The user will be created with the given, comma seperated, roles`,
	Run:   AddUser,
}

// Initialize the add command
func init() {
	rootCmd.AddCommand(addCommand)

	addCommand.Flags().StringP("name", "n", "", "Name of the new user")
	addCommand.Flags().StringP("email", "e", "", "Email address of the new user")
	addCommand.Flags().StringP("roles", "r", "Student", "Comma seperated roles of the new user")
}

// Adds a new user to the database
func AddUser(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	email, _ := cmd.Flags().GetString("email")
	roles, _ := cmd.Flags().GetString("roles")
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		configPath = "config.yaml"
	}

	// Load configuration
	config := c.FromFile(configPath)
	// Show the branding banner
	config.ShowBrand()

	// Initialize the database provider
	provider := db.NewProvider(config)
	// Initialize the user dao
	userDao := dao.NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))

	// Check if the email is empty
	if email == "" {
		config.Logger().Error("Email is required")
		return
	}

	// Check if the name is empty
	if name == "" {
		config.Logger().Error("Name is required")
		return
	}

	// Check if the mail is valid
	_, err = mail.ParseAddress(email)
	if err != nil {
		config.Logger().Error("Invalid email address")
		return
	}

	// Check if the mail is part of the allowed domains
	emailDomain := email[strings.Index(email, "@")+1:]
	if !slices.Contains(config.AllowedDomains, emailDomain) {
		config.Logger().Error("Email domain is not allowed to register users. Please add the domain to the allowed domains in the configuration file")
		return
	}

	// Split the roles
	roleList := strings.Split(roles, ",")

	// Create a list of database roles
	roleEntities := make([]*models.Role, 0)
	// Iterate over the roles
	for _, role := range roleList {
		// Get the role from the database
		roleEntity, err := userDao.GetRoleByClaim(role)
		// Check if the role exists
		if err != nil {
			config.Logger().Error(fmt.Sprintf("Role %s does not exist. Please create the role in the configuration first or run the backend at least once to migrate the database", role))
			return
		}
		// Add the role to the list
		roleEntities = append(roleEntities, roleEntity)
	}

	// Create the user
	user := models.User{
		Name:  name,
		Email: email,
		Roles: roleEntities,
	}

	existentUser, _ := userDao.GetByEmail(email)
	if existentUser != nil {
		config.Logger().Error("User already exists")
		return
	}

	// Add the user to the database
	_, daoErr := userDao.Create(user)
	if daoErr != nil {
		config.Logger().Error("Failed to create user")
		return
	}

	config.Logger().Info("User sucessfully created")

}
