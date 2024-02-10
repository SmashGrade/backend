package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type FieldRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewFieldRepository(provider db.Provider) *FieldRepository {
	return &FieldRepository{
		Provider: provider,
	}
}

func (r *FieldRepository) DeleteId(id uint) error {
	return r.Provider.DB().Delete(&models.Field{}, id).Error
}
