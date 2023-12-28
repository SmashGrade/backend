package provider

import (
	"log"

	"github.com/SmashGrade/backend/app/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultDbPath string = "data.db"
const autoMigrateAtConnect bool = true

type SqliteProvider struct {
	db *gorm.DB
}

// automigrates all entities
func (s *SqliteProvider) Migrate() {
	// Migrate all free standing structs first
	s.db.AutoMigrate(&entity.Curriculumtype{})
	s.db.AutoMigrate(&entity.Field{})
	s.db.AutoMigrate(&entity.Evaluationtype{})
	s.db.AutoMigrate(&entity.Examtype{})
	s.db.AutoMigrate(&entity.Role{})
	s.db.AutoMigrate(&entity.State{})
	s.db.AutoMigrate(&entity.Gradetype{})
	// Migrate secondary structs second
	s.db.AutoMigrate(&entity.Focus{})
	s.db.AutoMigrate(&entity.Module{})
	s.db.AutoMigrate(&entity.Curriculum{})
	s.db.AutoMigrate(&entity.Course{})
	s.db.AutoMigrate(&entity.User{})
	// Migrate convoluted stuff last
	s.db.AutoMigrate(&entity.SelectedCourse{})
	s.db.AutoMigrate(&entity.Exam{})
	s.db.AutoMigrate(&entity.ExamEvaluation{})
	s.db.AutoMigrate(&entity.Conversion{})
}

// connect to the defaultDbPath sqlite
func (s *SqliteProvider) Connect() {
	db, err := gorm.Open(sqlite.Open(defaultDbPath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	s.db = db

	if autoMigrateAtConnect {
		s.Migrate()
	}
}
