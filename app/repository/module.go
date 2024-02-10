package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type ModuleRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewModuleRepository(provider db.Provider) *ModuleRepository {
	return &ModuleRepository{
		Provider: provider,
	}
}

func (r *ModuleRepository) DeleteVersioned(id uint, version uint) error {
	return r.Provider.DB().Where("id = ? AND version = ?", id, version).Delete(&models.Module{}).Error
}

func (r *ModuleRepository) GetLatestId() (id uint, err error) {
	result := r.Provider.DB().Select("max(id) as id").First(&models.Module{}).Pluck("id", &id)
	err = result.Error
	return
}
