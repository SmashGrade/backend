package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type FocusRepository struct {
	*BaseRepository
}

func NewFocusRepository(provider db.Provider) *FocusRepository {
	return &FocusRepository{
		BaseRepository: NewBaseRepository(provider, models.Focus{}),
	}
}
