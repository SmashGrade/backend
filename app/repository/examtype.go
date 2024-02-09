package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type ExamtypeRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewExamtypeRepository(provider db.Provider) *ExamtypeRepository {
	return &ExamtypeRepository{
		Provider: provider,
	}
}

func (r *ExamtypeRepository) DeleteId(id uint) error {
	result := r.Provider.DB().Delete(&models.Examtype{}, id)
	return result.Error
}
