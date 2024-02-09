package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type CourseRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewCourseRepository(provider db.Provider) *CourseRepository {
	return &CourseRepository{
		Provider: provider,
	}
}

func (r *CourseRepository) DeleteVersioned(id, version uint) error {
	result := r.Provider.DB().Where("id = ? AND version = ?", id, version).Delete(&models.Course{})
	return result.Error
}

func (r *CourseRepository) GetLatestetId() (id uint, err error) {
	result := r.Provider.DB().Select("max(id) as id").First(&models.Course{}).Pluck("id", &id)
	err = result.Error
	return
}
