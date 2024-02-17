package repository

import (
	"reflect"
	"time"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"gorm.io/gorm/clause"
)

type ConversionRepository struct {
	*BaseRepository
}

func NewConversionRepository(provider db.Provider) *ConversionRepository {
	return &ConversionRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}

func (r *ConversionRepository) GetTimed(id uint, startDate time.Time, entity any) (any, error) {
	// Get tye of entity
	dtype := reflect.TypeOf(entity)
	// Create a new instance of the entity type
	newEntity := reflect.New(dtype).Interface()

	result := r.Provider.DB().Preload(clause.Associations).
		Where("id = ? AND ee_selected_course_class_startyear = ?", id, startDate).
		First(newEntity)
	if result.Error != nil {
		return models.Conversion{}, result.Error
	}
	return newEntity, nil
}

func (r *ConversionRepository) DeleteTimed(id uint, startDate time.Time, entity any) error {
	// Get tye of entity
	dtype := reflect.TypeOf(entity)
	// Create a new instance of the entity type
	newEntity := reflect.New(dtype).Interface()

	return r.Provider.DB().
		Where("id = ? AND ee_selected_course_class_startyear = ?", id, startDate).
		Delete(newEntity).Error
}
