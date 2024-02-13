package repository

import (
	"time"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"gorm.io/gorm/clause"
)

type CurriculumRepository struct {
	*BaseRepository
}

func NewCurriculumRepository(provider db.Provider) *CurriculumRepository {
	return &CurriculumRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}

func (r *CurriculumRepository) DeleteTimed(id uint, startDate time.Time) error {
	return r.Provider.DB().
		Where("id = ? AND start_validity = ?", id, startDate).
		Delete(&models.Curriculum{}).Error
}

func (r *CurriculumRepository) GetTimed(id uint, startValidity time.Time) (entity models.Curriculum, err error) {
	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND start_validity = ?", id, startValidity).First(&entity)
	err = result.Error
	return
}
