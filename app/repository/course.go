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
	if result.Error != nil {
		return result.Error
	}
	return nil
}
