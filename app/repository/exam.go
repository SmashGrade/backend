package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type ExamRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewExamRepository(provider db.Provider) *ExamRepository {
	return &ExamRepository{
		Provider: provider,
	}
}

func (r *ExamRepository) DeleteId(id uint) error {
	result := r.Provider.DB().Delete(&models.Exam{}, id)
	return result.Error
}
