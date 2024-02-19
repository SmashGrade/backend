package repository

import (
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
		BaseRepository: NewBaseRepository(provider, models.Conversion{}),
	}
}

func (r *ConversionRepository) GetTimed(id uint, startDate time.Time) (any, error) {
	newEntity := r.getInterface()

	result := r.Provider.DB().Preload(clause.Associations).
		Where("id = ? AND ee_selected_course_class_startyear = ?", id, startDate).
		First(newEntity)
	if result.Error != nil {
		return models.Conversion{}, result.Error
	}
	return newEntity, nil
}

func (r *ConversionRepository) DeleteTimed(id uint, startDate time.Time) error {
	newEntity := r.getInterface()

	return r.Provider.DB().
		Where("id = ? AND ee_selected_course_class_startyear = ?", id, startDate).
		Delete(newEntity).Error
}
