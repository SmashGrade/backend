package db

import (
	c "github.com/SmashGrade/backend/app/config"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"gorm.io/gorm"
)

type Provider interface {
	DB() *gorm.DB                          // Returns the database connection
	Config() *c.APIConfig                  // Returns the API configuration
	Connect(connectionString string) error // Connect to the database
	IsConnected() bool                     // Returns true if the provider is connected to the database
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
func (s *SQLiteProvider) Connect(connectionString string) error {
	return nil
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
	}
	return nil
}
