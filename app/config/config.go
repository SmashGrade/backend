package config

import "fmt"

// APIConfig is used to define the configuration of the API
type APIConfig struct {
	Host        string // Host is the address of the server
	Port        int    // Port is the port of the server
	AutoMigrate bool   // AutoMigrate is a flag to determine if the database should be migrated automatically

}

// Returns a new configuration with default values
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		Host:        "0.0.0.0",
		Port:        9000,
		AutoMigrate: true,
	}
}

// Returns the server address as a string
func (c *APIConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
