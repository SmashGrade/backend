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
		BaseRepository: NewBaseRepository(provider, models.Role{}),
	}
}
