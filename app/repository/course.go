package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type CourseRepository struct {
	*BaseRepository
}

func NewCourseRepository(provider db.Provider) *CourseRepository {
	return &CourseRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}

func (r *CourseRepository) DeleteVersioned(id, version uint) error {
	result := r.Provider.DB().Where("id = ? AND version = ?", id, version).Delete(&models.Course{})
	return result.Error
}

func (r *CourseRepository) GetLatestId() (id uint, err error) {
	result := r.Provider.DB().Select("max(id) as id").First(&models.Course{}).Pluck("id", &id)
	err = result.Error
	return
}

// returns highest versioned entity
func (r *CourseRepository) GetLatestVersioned(id uint) (entity any, err error) {
	err = r.Provider.DB().Where("id = ?", id).Order("version desc").First(entity).Error
	return
}
