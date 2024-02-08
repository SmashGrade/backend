package config

// Package config is used to define the configuration of the application
type APIConfig struct {
	Host        string // Host is the address of the server
	Port        int    // Port is the port of the server
	AutoMigrate bool   // AutoMigrate is a flag to determine if the database should be migrated automatically

}

// NewAPIConfig returns a new APIConfig
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		Host:        "0.0.0.0",
		Port:        9000,
		AutoMigrate: true,
	}
}
