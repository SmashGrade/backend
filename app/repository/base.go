package repository

import (
	"time"

	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(entity any) error
	Update(entity any) error
	Find(entity any) (entities []any, err error)
	GetAll() (entites []any, err error)
}

// Repository methods for models with only an id
type IdRepository interface {
	GetId(id uint) (entity any, err error)
	DeleteId(id uint) error
}

// Repository methods for models with id and version
type VersionedRepository interface {
	GetVersioned(id, version uint) (entity any, err error)
	DeleteVersioned(id, version uint) error
}

// Repository methods for models with id and start date
type TimedRepository interface {
	GetTimed(id uint, startDate time.Time) (entity any, err error)
	DeleteTimed(id uint, startDate time.Time) error
}

// BaseRepository is a base repository
// that contains the database connection and CRUD operations
type BaseRepository struct {
	Provider *db.BaseProvider
}

// Constructor for BaseRepository
func NewBaseRepository(provider *db.BaseProvider) *BaseRepository {
	return &BaseRepository{
		Provider: provider,
	}
}

// Example functions
// TODO: Please implement them in the actual repository concrete for the model
func (r *BaseRepository) Create(entity any) error {
	return e.ErrNotImplemented
}

func (r *BaseRepository) Update(entity any) error {
	return e.ErrNotImplemented
}

func (r *BaseRepository) Find(entity any) (entities []any, err error) {
	return nil, e.ErrNotImplemented
}

func (r *BaseRepository) GetAll() (entities []any, err error) {
	return nil, e.ErrNotImplemented
}

// For versions and timed repositories you also need to implement the following methods
func (r *BaseRepository) Get(id uint, version uint) (entity interface{}, err error) {
	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND version = ?", id, version).First(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

// func (r *BaseRepository) Delete(id uint, version uint) error

// func (r *BaseRepository) Get(id uint, startDate time.Time) (entity any, err error)
// func (r *BaseRepository) Delete(id uint, startDate time.Time) error
