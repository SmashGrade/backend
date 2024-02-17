package repository

import (
	"github.com/SmashGrade/backend/app/db"
)

type FocusRepository struct {
	*BaseRepository
}

func NewFocusRepository(provider db.Provider) *FocusRepository {
	return &FocusRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}
