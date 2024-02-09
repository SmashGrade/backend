package repository

import "gorm.io/gorm"

// BaseRepository is a base repository
// that contains the database connection and CRUD operations
type BaseRepository struct {
	Db *gorm.DB
}

// Constructor for BaseRepository
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		Db: db,
	}
}
