package repository

import (
	"time"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"gorm.io/gorm/clause"
)

type ConversionRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewConversionRepository(provider db.Provider) *ConversionRepository {
	return &ConversionRepository{
		Provider: provider,
	}
}

func (r *ConversionRepository) GetTimed(id uint, startDate time.Time) (entity models.Conversion, err error) {
	result := r.Provider.DB().Preload(clause.Associations).
		Where("id = ? AND ee_selected_course_class_startyear = ?", id, startDate).
		First(&entity)
	err = result.Error
	return
}

func (r *ConversionRepository) DeleteTimed(id uint, startDate time.Time) error {
	return r.Provider.DB().
		Where("id = ? AND ee_selected_course_class_startyear = ?", id, startDate).
		Delete(&models.Conversion{}).Error
}
