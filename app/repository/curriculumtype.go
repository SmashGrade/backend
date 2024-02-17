package repository

import (
	"github.com/SmashGrade/backend/app/db"
)

type CurriculumtypeRepository struct {
	*BaseRepository
}

func NewCurriculumtypeRepository(provider db.Provider) *CurriculumtypeRepository {
	return &CurriculumtypeRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}
