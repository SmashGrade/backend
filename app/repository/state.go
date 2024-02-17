package repository

import (
	"github.com/SmashGrade/backend/app/db"
)

type StateRepository struct {
	*BaseRepository
}

func NewStateRepository(provider db.Provider) *StateRepository {
	return &StateRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}
