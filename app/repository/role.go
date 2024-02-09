package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type RoleRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewRoleRepository(provider db.Provider) *RoleRepository {
	return &RoleRepository{
		Provider: provider,
	}
}

func (r *RoleRepository) DeleteId(id uint) error {
	result := r.Provider.DB().Delete(&models.Role{}, id)
	return result.Error
}
