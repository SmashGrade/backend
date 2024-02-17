package repository

import (
	"reflect"
	"time"

	"github.com/SmashGrade/backend/app/db"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(entity any) (any, error)
	Update(entity any) error
	Find(entity any) (any, error)
	GetAll() (any, error)
}

// Repository methods for models with only an id
type IdRepository interface {
	GetId(id uint) (any, error)
	DeleteId(id uint) error
}

// Repository methods for models with id and version
type VersionedRepository interface {
	GetVersioned(id, version uint) (any, error)
	DeleteVersioned(id, version uint) error
	GetLatestVersioned(id uint) (any, error)
}

// Repository methods for models with id and start date
type TimedRepository interface {
	GetTimed(id uint, startDate time.Time) (any, error)
	DeleteTimed(id uint, startDate time.Time) error
}

// BaseRepository is a base repository
// that contains the database connection and CRUD operations
type BaseRepository struct {
	Provider db.Provider
	Type     any
}

// Constructor for BaseRepository
func NewBaseRepository(provider db.Provider, entity any) *BaseRepository {
	return &BaseRepository{
		Provider: provider,
		Type:     entity,
	}
}

// Example functions
// TODO: Please implement them in the actual repository concrete for the model
func (r *BaseRepository) Create(entity any) (any, error) {
	result := r.Provider.DB().Create(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (r *BaseRepository) Update(entity any) error {
	return r.Provider.DB().Updates(entity).Error
}

func (r *BaseRepository) Find(entity any) (any, error) {
	// Get type of entity
	dtype := reflect.TypeOf(entity)
	// Create a new Slice of the entity type
	entities := reflect.New(reflect.SliceOf(dtype)).Elem().Interface()

	result := r.Provider.DB().Preload(clause.Associations).Where(&entity).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

func (r *BaseRepository) GetAll() (any, error) {
	// Get type of entity
	dtype := reflect.TypeOf(r.Type)
	// Create a new Slice of the entity type
	entities := reflect.New(reflect.SliceOf(dtype)).Elem().Interface()

	result := r.Provider.DB().Preload(clause.Associations).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

func (r *BaseRepository) GetId(id uint) (any, error) {
	// Get tye of entity
	dtype := reflect.TypeOf(r.Type)
	// Create a new instance of the entity type
	newEntity := reflect.New(dtype).Interface()

	result := r.Provider.DB().Preload(clause.Associations).First(newEntity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

func (r *BaseRepository) GetVersioned(id, version uint) (any, error) {
	// Get tye of entity
	dtype := reflect.TypeOf(r.Type)
	// Create a new instance of the entity type
	newEntity := reflect.New(dtype).Interface()

	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND version = ?", id, version).First(newEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

func (r *BaseRepository) GetTimed(id uint, startDate time.Time) (any, error) {
	// Get tye of entity
	dtype := reflect.TypeOf(r.Type)
	// Create a new instance of the entity type
	newEntity := reflect.New(dtype).Interface()

	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND startyear = ?", id, startDate).First(newEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

func (r *BaseRepository) DeleteVersioned(id, version uint) error {
	// Get tye of entity
	dtype := reflect.TypeOf(r.Type)
	// Create a new instance of the entity type
	newEntity := reflect.New(dtype).Interface()

	return r.Provider.DB().Where("id = ? AND version = ?", id, version).Delete(newEntity).Error
}

func (r *BaseRepository) GetLatestVersioned(id uint) (any, error) {
	// Get tye of entity
	dtype := reflect.TypeOf(r.Type)
	// Create a new instance of the entity type
	newEntity := reflect.New(dtype).Interface()

	result := r.Provider.DB().Where("id = ?", id).Order("version desc").First(newEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

func (r *BaseRepository) DeleteId(id uint) error {
	// Get tye of entity
	dtype := reflect.TypeOf(r.Type)
	// Create a new instance of the entity type
	newEntity := reflect.New(dtype).Interface()

	return r.Provider.DB().Delete(newEntity, id).Error
}
