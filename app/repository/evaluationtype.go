package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type EvaluationtypeRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewEvaluationtypeRepository(provider db.Provider) *EvaluationtypeRepository {
	return &EvaluationtypeRepository{
		Provider: provider,
	}
}

func (r *EvaluationtypeRepository) DeleteId(id uint) error {
	result := r.Provider.DB().Delete(&models.Evaluationtype{}, id)
	return result.Error
}
