package repository

import (
	"time"

	"github.com/SmashGrade/backend/app/db"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(entity any) (returnEntity any, err error)
	Update(entity any) (err error)
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
	GetLatestVersioned(id uint) (entity any, err error)
}

// Repository methods for models with id and start date
type TimedRepository interface {
	GetTimed(id uint, startDate time.Time) (entity any, err error)
	DeleteTimed(id uint, startDate time.Time) error
}

// BaseRepository is a base repository
// that contains the database connection and CRUD operations
type BaseRepository struct {
	Provider db.Provider
}

// Constructor for BaseRepository
func NewBaseRepository(provider db.Provider) *BaseRepository {
	return &BaseRepository{
		Provider: provider,
	}
}

// Example functions
// TODO: Please implement them in the actual repository concrete for the model
func (r *BaseRepository) Create(entity any) (returnEntity any, err error) {
	result := r.Provider.DB().Create(entity)
	err = result.Error
	returnEntity = entity
	return
}

func (r *BaseRepository) Update(entity any) error {
	return r.Provider.DB().Model(&entity).Updates(&entity).Error
}

func (r *BaseRepository) Find(entity any) (entities []any, err error) {
	result := r.Provider.DB().Preload(clause.Associations).Where(&entity).Find(&entities)
	err = result.Error
	return
}

func (r *BaseRepository) GetAll() (entities []any, err error) {
	result := r.Provider.DB().Preload(clause.Associations).Find(&entities)
	err = result.Error
	return
}

func (r *BaseRepository) GetId(id uint) (entity any, err error) {
	result := r.Provider.DB().Preload(clause.Associations).First(&entity, id)
	err = result.Error
	return
}

func (r *BaseRepository) GetVersioned(id uint, version uint) (entity any, err error) {
	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND version = ?", id, version).First(&entity)
	err = result.Error
	return
}

func (r *BaseRepository) GetTimed(id uint, startDate time.Time) (entity any, err error) {
	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND startyear = ?", id, startDate).First(&entity)
	err = result.Error
	return
}
