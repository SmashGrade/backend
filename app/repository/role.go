package repository

import (
	"github.com/SmashGrade/backend/app/db"
)

type RoleRepository struct {
	*BaseRepository
}

func NewRoleRepository(provider db.Provider) *RoleRepository {
	return &RoleRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}
