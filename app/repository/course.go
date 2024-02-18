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
