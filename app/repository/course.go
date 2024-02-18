package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/google/uuid"
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
func (r *CourseRepository) GetLatestVersion(id uuid.UUID) (uint, error) {
	retCourse := &models.Course{}

	result := r.Provider.DB().Where("id = ?", id).Order("version desc").First(retCourse)
	if result.Error != nil {
		return 0, result.Error
	}
	return retCourse.Version, nil
}

// Returns next free version
func (r *CourseRepository) GetNextVersion(id uuid.UUID) (uint, error) {
	currentId, err := r.GetLatestVersion(id)
	if err != nil {
		currentId += 1
	}
	return currentId, err
}
