package repository

import (
	"github.com/SmashGrade/backend/app/db"
)

type FieldRepository struct {
	*BaseRepository
}

func NewFieldRepository(provider db.Provider) *FieldRepository {
	return &FieldRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}
