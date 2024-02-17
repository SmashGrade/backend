package repository

import (
	"github.com/SmashGrade/backend/app/db"
)

type StudyStageRepository struct {
	*BaseRepository
}

func NewStudyStageRepository(provider db.Provider) *StudyStageRepository {
	return &StudyStageRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}
