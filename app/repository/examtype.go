package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type ExamtypeRepository struct {
	*BaseRepository
}

func NewExamtypeRepository(provider db.Provider) *ExamtypeRepository {
	return &ExamtypeRepository{
		BaseRepository: NewBaseRepository(provider, models.Examtype{}),
	}
}
