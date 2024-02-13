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
		BaseRepository: NewBaseRepository(provider),
	}
}

func (r *FocusRepository) DeleteId(id uint) error {
	return r.Provider.DB().Delete(&models.Focus{}, id).Error
}
