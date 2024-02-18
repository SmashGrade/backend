package repository

import (
	"reflect"
	"time"

	"github.com/SmashGrade/backend/app/db"
	"github.com/google/uuid"
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
	GetVersioned(id uuid.UUID, version uint) (any, error)
	DeleteVersioned(id uuid.UUID, version uint) error
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

// Return the slice of the Type in the BaseRepository
func (r *BaseRepository) getSliceInterface() any {
	// Get the type of Type
	dtype := reflect.TypeOf(r.Type)
	// Create a new Slice of the entity type
	return reflect.New(reflect.SliceOf(dtype)).Elem().Interface()
}

// Return the interface of the Type in the BaseRepository
func (r *BaseRepository) getInterface() any {
	// Get the type of Type
	dtype := reflect.TypeOf(r.Type)
	// Create a new instance of the entity type
	return reflect.New(dtype).Interface()
}

/*
Creates new Entity on the DB.

WARNING: Expect Pointer or else it will panic

Usage (example with models.Course):

	newCourse, err := repository.Create(&course)
*/
func (r *BaseRepository) Create(entity any) (any, error) {
	result := r.Provider.DB().Create(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

/*
Updates Entity on the DB.

Usage (example with models.Course):

	err := repository.Update(&course)
*/
func (r *BaseRepository) Update(entity any) error {
	return r.Provider.DB().Updates(entity).Error
}

/*
Get list of entitys providing a entity

Usage (example with models.Course):

	res, err := repository.Find(course)
	courses := res.([]models.Course)
*/
func (r *BaseRepository) Find(entity any) (any, error) {
	entities := r.getSliceInterface()

	result := r.Provider.DB().Preload(clause.Associations).Where(&entity).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

/*
Gets list of entity from DB

Usage (example with models.Course):

	res, err := repository.GetAll()
	courses := res.([]models.Course)
*/
func (r *BaseRepository) GetAll() (any, error) {
	entities := r.getSliceInterface()

	result := r.Provider.DB().Preload(clause.Associations).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

/*
Get entity by providing the id

Usage (example with models.Field)

	res, err := repository.GetId(3)
	field := res.(*models.Field)
*/
func (r *BaseRepository) GetId(id uint) (any, error) {
	newEntity := r.getInterface()

	result := r.Provider.DB().Preload(clause.Associations).First(newEntity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

/*
Get entity by providing the id and version

Usage (example with models.Course)

	res, err := repository.GetVersioned(1, 2)
	course := result.(*models.Course)
*/
func (r *BaseRepository) GetVersioned(id uuid.UUID, version uint) (any, error) {
	newEntity := r.getInterface()

	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND version = ?", id, version).First(newEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

/*
Get entity by providing the id and startdate

Usage (example with models.Conversion):

	res, err := repository.GetTimed(1, time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC))
	conversion := res.(*models.Conversion)
*/
func (r *BaseRepository) GetTimed(id uint, startDate time.Time) (any, error) {
	newEntity := r.getInterface()

	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND startyear = ?", id, startDate).First(newEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

/*
Delete entity by providing the id and version

Usage (example with models.Course):

	err := respository.DeleteVersioned(1, 2)
*/
func (r *BaseRepository) DeleteVersioned(id uuid.UUID, version uint) error {
	newEntity := r.getInterface()

	return r.Provider.DB().Where("id = ? AND version = ?", id, version).Delete(newEntity).Error
}

/*
Get entity with the highst version, providing the id

Usage (example with models.Course):

	res, err := repository.GetLatestVersioned(1)
	course := res.(*models.Course)
*/
func (r *BaseRepository) GetLatestVersioned(id uuid.UUID) (any, error) {
	newEntity := r.getInterface()

	result := r.Provider.DB().Where("id = ?", id).Order("version desc").First(newEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

/*
Delete entity by providing the id

Usage (example with models.Field):

	err := respository.DeleteId(1)
*/
func (r *BaseRepository) DeleteId(id uint) error {
	newEntity := r.getInterface()

	return r.Provider.DB().Delete(newEntity, id).Error
}
