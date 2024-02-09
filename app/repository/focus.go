package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type FocusRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewFocusRepository(provider db.Provider) *FocusRepository {
	return &FocusRepository{
		Provider: provider,
	}
}

func (r *FocusRepository) DeleteId(id uint) error {
	result := r.Provider.DB().Delete(&models.Focus{}, id)
	return result.Error
}
