package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type ModuleRepository struct {
	*BaseRepository
}

func NewModuleRepository(provider db.Provider) *ModuleRepository {
	return &ModuleRepository{
		BaseRepository: NewBaseRepository(provider, models.Module{}),
	}
}
