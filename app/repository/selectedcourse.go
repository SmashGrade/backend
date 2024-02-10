package repository

import (
	"time"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type SelectedCourseRepository struct {
	Provider db.Provider
	*BaseRepository
}

func NewSelectedCourseRepository(provider db.Provider) *SelectedCourseRepository {
	return &SelectedCourseRepository{
		Provider: provider,
	}
}

func (r *SelectedCourseRepository) GetSelectedCourse(
	userId, courseId, courseVersion uint, classStartyear time.Time) (selectedCourse models.SelectedCourse, err error) {
	result := r.Provider.DB().
		Where("user_id = ? AND course_id = ? AND course_version = ? AND class_startyear = ?",
			userId,
			courseId,
			courseVersion,
			classStartyear).First(&selectedCourse)
	err = result.Error
	return
}

func (r *SelectedCourseRepository) DeleteSelectedCourse(
	userId, courseId, courseVersion uint, classStartyear time.Time) error {
	return r.Provider.DB().
		Where("user_id = ? AND course_id = ? AND course_version = ? AND class_startyear = ?",
			userId,
			courseId,
			courseVersion,
			classStartyear).Delete(&models.SelectedCourse{}).Error
}
