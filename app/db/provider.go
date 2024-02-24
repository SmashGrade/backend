package db

import (
	"fmt"
	"strings"

	c "github.com/SmashGrade/backend/app/config"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Provider interface {
	DB() *gorm.DB         // Returns the database connection
	Config() *c.APIConfig // Returns the API configuration
	Connect() error       // Connect to the database
	IsConnected() bool    // Returns true if the provider is connected to the database
}

type BaseProvider struct {
	config      *c.APIConfig // API configuration
	Db          *gorm.DB     // Database connection
	isConnected bool         // True if the provider is connected to the database
}

// Returns the database connection
func (p *BaseProvider) DB() *gorm.DB {
	return p.Db
}

// Connect to the database
func (p *BaseProvider) Connect(connectionString string) error {
	// Not implemented in the base provider
	return e.ErrNotImplemented
}

// Returns true if the provider is connected to the database
func (p *BaseProvider) IsConnected() bool {
	return p.isConnected
}

// Returns the API configuration
func (p BaseProvider) Config() *c.APIConfig {
	return p.config
}

type SQLiteProvider struct {
	BaseProvider
}

// Connect to the database
// Expecting a connection string as sqlite://{path}
func (s *SQLiteProvider) Connect() error {
	// If already connected, return nil
	if s.IsConnected() {
		return nil
	}
	// Fix the connection string, remove the sqlite:// prefix
	connectionStr := strings.TrimPrefix(s.config.DBConnectionStr, "sqlite://")
	// Open the database connection
	db, err := gorm.Open(sqlite.Open(connectionStr))
	if err != nil {
		return err
	}
	// Assign the database connection
	s.Db = db
	s.isConnected = true
	s.config.Logger().Info(fmt.Sprintf("Connected to database %s", s.config.DBConnectionStr))
	// Migrate the database
	err = Migrate(s)
	if err != nil {
		return err
	}
	return nil
}

func NewSQLiteProvider(config *c.APIConfig) *SQLiteProvider {
	return &SQLiteProvider{
		BaseProvider{
			config: config,
		},
	}
}

// Migrates the database with all existing models
func Migrate(p Provider) error {
	models := []interface{}{
		// Add all models here
		models.Curriculumtype{},
		models.Field{},
		models.Evaluationtype{},
		models.Examtype{},
		models.Role{},
		models.State{},
		models.Gradetype{},
		models.StudyStage{},
		models.Focus{},
		models.Module{},
		models.Curriculum{},
		models.Course{},
		models.User{},
		models.SelectedCourse{},
		models.Exam{},
		models.ExamEvaluation{},
		models.Conversion{},
	}
	// Migrate all models if autoMigrateAtConnect is true
	if p.Config().AutoMigrate {
		for _, model := range models {
			err := p.DB().AutoMigrate(model)
			if err != nil {
				return err
			}
		}
		p.Config().Logger().Info(fmt.Sprintf("Database %s migrated successfully", p.Config().DBConnectionStr))
	} else {
		p.Config().Logger().Warn("Database migration is disabled")
	}
	return nil
}

type PostgresProvider struct {
	BaseProvider
}

// Connect to the database
func (p *PostgresProvider) Connect() error {
	// Not implemented
	return e.ErrNotImplemented
}

// Returns a new Postgres provider
func NewPostgresProvider(config *c.APIConfig) *PostgresProvider {
	return &PostgresProvider{
		BaseProvider{
			config: config,
		},
	}
}

type MySQLProvider struct {
	BaseProvider
}

// Connect to the database
func (m *MySQLProvider) Connect() error {
	// Not implemented
	return e.ErrNotImplemented
}

// Returns a new MySQL provider
func NewMySQLProvider(config *c.APIConfig) *MySQLProvider {
	return &MySQLProvider{
		BaseProvider{
			config: config,
		},
	}
}

// Returns a new provider based on the connection string
func NewProvider(config *c.APIConfig) Provider {
	var provider Provider
	switch true {
	case strings.HasPrefix(config.DBConnectionStr, "sqlite://"):
		provider = NewSQLiteProvider(config)
	case strings.HasPrefix(config.DBConnectionStr, "postgres://"):
		provider = NewPostgresProvider(config)
	case strings.HasPrefix(config.DBConnectionStr, "mysql://"):
		provider = NewMySQLProvider(config)
	default:
		provider = nil
	}
	// Check connection and migrate database
	if config.Connect {
		err := provider.Connect()
		if err != nil {
			config.Logger().Fatal(fmt.Sprintf("Failed to connect to database %s: %s", config.DBConnectionStr, err))
		}
	}
	if config.MockData {
		prefillMockDB(provider)
	}
	return provider
}

// Returns an in memory provider for mocking and testing
func NewMockProvider() Provider {
	// Override provider connection string with sqlite memory
	var mockConfig *c.APIConfig = c.NewAPIConfig()
	// Override connection string
	mockConfig.DBConnectionStr = "sqlite://:memory:"
	// Ensure that DB in memory is always migrated
	mockConfig.AutoMigrate = true
	provider := NewProvider(mockConfig)

	return provider
}

// Returns an in memory provider for mocking and testing with prefilled data
func NewPrefilledMockProvider() Provider {

	// Override provider connection string with sqlite memory
	var mockConfig *c.APIConfig = c.NewAPIConfig()
	mockConfig.MockData = true
	// Override connection string
	mockConfig.DBConnectionStr = "sqlite://:memory:"
	// Ensure that DB in memory is always migrated
	mockConfig.AutoMigrate = true
	provider := NewProvider(mockConfig)

	return provider
}
