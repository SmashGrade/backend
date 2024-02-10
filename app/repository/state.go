package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type StateRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewStateRepository(provider db.Provider) *StateRepository {
	return &StateRepository{
		Provider: provider,
	}
}

func (r *StateRepository) DeleteId(id uint) error {
	return r.Provider.DB().Delete(&models.State{}, id).Error
}
