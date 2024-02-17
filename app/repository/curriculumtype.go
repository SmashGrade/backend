package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type CurriculumtypeRepository struct {
	*BaseRepository
}

func NewCurriculumtypeRepository(provider db.Provider) *CurriculumtypeRepository {
	return &CurriculumtypeRepository{
		BaseRepository: NewBaseRepository(provider, models.Curriculum{}),
	}
}
