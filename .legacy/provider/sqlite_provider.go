package provider

import (
	"log"

	"github.com/SmashGrade/backend/legacy/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultDbPath string = "data.db"
const autoMigrateAtConnect bool = true

type SqliteProvider struct {
	Db *gorm.DB
}

// automigrates all entities
func (s *SqliteProvider) migrate() {
	// Migrate all free standing structs first
	s.Db.AutoMigrate(&entity.Curriculumtype{})
	s.Db.AutoMigrate(&entity.Field{})
	s.Db.AutoMigrate(&entity.Evaluationtype{})
	s.Db.AutoMigrate(&entity.Examtype{})
	s.Db.AutoMigrate(&entity.Role{})
	s.Db.AutoMigrate(&entity.State{})
	s.Db.AutoMigrate(&entity.Gradetype{})
	s.Db.AutoMigrate(&entity.StudyStage{})
	// Migrate secondary structs second
	s.Db.AutoMigrate(&entity.Focus{})
	s.Db.AutoMigrate(&entity.Module{})
	s.Db.AutoMigrate(&entity.Curriculum{})
	s.Db.AutoMigrate(&entity.Course{})
	s.Db.AutoMigrate(&entity.User{})
	// Migrate convoluted stuff last
	s.Db.AutoMigrate(&entity.SelectedCourse{})
	s.Db.AutoMigrate(&entity.Exam{})
	s.Db.AutoMigrate(&entity.ExamEvaluation{})
	s.Db.AutoMigrate(&entity.Conversion{})
}

// connect to the defaultDbPath sqlite
func (s *SqliteProvider) Connect() {
	db, err := gorm.Open(sqlite.Open(defaultDbPath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	s.Db = db

	if autoMigrateAtConnect {
		s.migrate()
	}
}
