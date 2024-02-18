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
		BaseRepository: NewBaseRepository(provider, models.Course{}),
	}
}

// Returns currently highest used version
func (r *CourseRepository) GetLatestVersion(id uint) (uint, error) {

	ret, err := r.GetLatestVersioned(id)
	if err != nil {
		return 0, err
	}

	retCourse := ret.(*models.Module)
	return retCourse.Version, nil
}

// Returns next free version
func (r *CourseRepository) GetNextVersion(id uint) (uint, error) {
	currentId, err := r.GetLatestVersion(id)
	if err == nil {
		currentId += 1
	}
	return currentId, err
}
