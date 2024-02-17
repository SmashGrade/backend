package repository

import (
	"github.com/SmashGrade/backend/app/db"
)

type CourseRepository struct {
	*BaseRepository
}

func NewCourseRepository(provider db.Provider) *CourseRepository {
	return &CourseRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}
