package db

import "gorm.io/gorm"

type Provider interface {
	DB() *gorm.DB
	Connect(connectionString string) error
	IsConnected() bool
}

type BaseProvider struct {
	Db          *gorm.DB
	isConnected bool
}

// Returns the database connection
func (p *BaseProvider) DB() *gorm.DB {
	return p.Db
}

// Connect to the database
func (p *BaseProvider) Connect(connectionString string) error {
	if !p.IsConnected() {

	}
}

// Returns true if the provider is connected to the database
func (p *BaseProvider) IsConnected() bool {
	return p.isConnected
}

type SQLiteProvider struct {
	BaseProvider
}

// Connect to the database
// Expecting a connection string as sqlite://{path}
func (s *SQLiteProvider) Connect(connectionString string) error {
	return nil
}

// Migrates the database with all existing models
func Migrate(p Provider) error {

}
