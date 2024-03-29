package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type EvaluationtypeRepository struct {
	*BaseRepository
}

func NewEvaluationtypeRepository(provider db.Provider) *EvaluationtypeRepository {
	return &EvaluationtypeRepository{
		BaseRepository: NewBaseRepository(provider, models.Evaluationtype{}),
	}
}
