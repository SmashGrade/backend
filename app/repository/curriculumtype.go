package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type CurriculumtypeRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewCurriculumtypeRepository(provider db.Provider) *CurriculumtypeRepository {
	return &CurriculumtypeRepository{
		Provider: provider,
	}
}

func (r *CurriculumtypeRepository) DeleteId(id uint) error {
	result := r.Provider.DB().Delete(&models.Curriculumtype{}, id)
	return result.Error
}
