package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type GradetypeRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewGradetypeRepository(provider db.Provider) *GradetypeRepository {
	return &GradetypeRepository{
		Provider: provider,
	}
}

func (r *GradetypeRepository) DeleteId(id uint) error {
	return r.Provider.DB().Delete(&models.Gradetype{}, id).Error
}
