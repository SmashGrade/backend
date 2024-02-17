package repository

import (
	"github.com/SmashGrade/backend/app/db"
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository(provider db.Provider) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}
