package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type CurriculumRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewCurriculumRepository(provider db.Provider) *CurriculumRepository {
	return &CurriculumRepository{
		Provider: provider,
	}
}

func (r *CurriculumRepository) Delete(id uint) error {
	result := r.Provider.DB().Delete(&models.Curriculum{}, id)
	return result.Error
}
