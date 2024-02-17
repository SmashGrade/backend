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

func (r *ModuleRepository) GetLatestId() (id uint, err error) {
	result := r.Provider.DB().Select("max(id) as id").First(&models.Module{}).Pluck("id", &id)
	err = result.Error
	return
}
