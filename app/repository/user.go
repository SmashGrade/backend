package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type UserRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewUserRepository(provider db.Provider) *UserRepository {
	return &UserRepository{
		Provider: provider,
	}
}

func (r *UserRepository) DeleteId(id uint) error {
	return r.Provider.DB().Delete(&models.User{}, id).Error
}
