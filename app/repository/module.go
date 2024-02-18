package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/google/uuid"
)

type ModuleRepository struct {
	*BaseRepository
}

func NewModuleRepository(provider db.Provider) *ModuleRepository {
	return &ModuleRepository{
		BaseRepository: NewBaseRepository(provider, models.Module{}),
	}
}

// Returns currently highest used version
func (r *ModuleRepository) GetLatestVersion(id uuid.UUID) (uint, error) {

	ret, err := r.GetLatestVersioned(id)
	if err != nil {
		return 0, err
	}

	retCourse := ret.(*models.Module)
	return retCourse.Version, nil
}

// Returns next free version
func (r *ModuleRepository) GetNextVersion(id uuid.UUID) (uint, error) {
	currentId, err := r.GetLatestVersion(id)
	if err == nil {
		currentId += 1
	}
	return currentId, err
}
