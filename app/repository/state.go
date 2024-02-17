package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type StateRepository struct {
	*BaseRepository
}

func NewStateRepository(provider db.Provider) *StateRepository {
	return &StateRepository{
		BaseRepository: NewBaseRepository(provider, models.State{}),
	}
}
