package provider

import (
	"log"

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
	//db.AutoMigrate(&entity.Whatever)
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
