package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type RoleRepository struct {
	*BaseRepository
}

func NewRoleRepository(provider db.Provider) *RoleRepository {
	return &RoleRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}

func (r *RoleRepository) DeleteId(id uint) error {
	return r.Provider.DB().Delete(&models.Role{}, id).Error
}
