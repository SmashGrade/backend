package repository

import (
	"time"

	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
)

type Repository interface {
	Create(entity any) error
	Update(entity any) error
	Find(entity any) (entities []any, err error)
	GetAll() (entites []any, err error)
}

// Repository methods for models with only an id
type IdRepository interface {
	Get(id uint) (entity any, err error)
	Delete(id uint) error
}

// Repository methods for models with id and version
type VersionedRepository interface {
	Get(id, version uint) (entity any, err error)
	Delete(id, version uint) error
}

// Repository methods for models with id and start date
type TimedRepository interface {
	Get(id uint, startDate time.Time) (entity any, err error)
	Delete(id uint, startDate time.Time) error
}

// BaseRepository is a base repository
// that contains the database connection and CRUD operations
type BaseRepository struct {
	Provider *db.Provider
}

// Constructor for BaseRepository
func NewBaseRepository(provider *db.Provider) *BaseRepository {
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
// func (r *BaseRepository) Get(id uint, version uint) (entity any, err error)
// func (r *BaseRepository) Delete(id uint, version uint) error

// func (r *BaseRepository) Get(id uint, startDate time.Time) (entity any, err error)
// func (r *BaseRepository) Delete(id uint, startDate time.Time) error
