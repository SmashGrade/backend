package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type FieldRepository struct {
	*BaseRepository
}

func NewFieldRepository(provider db.Provider) *FieldRepository {
	return &FieldRepository{
		BaseRepository: NewBaseRepository(provider, models.Field{}),
	}
}
