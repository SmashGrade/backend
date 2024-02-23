package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type ExamRepository struct {
	*BaseRepository
}

func NewExamRepository(provider db.Provider) *ExamRepository {
	return &ExamRepository{
		BaseRepository: NewBaseRepository(provider, models.Exam{}),
	}
}
