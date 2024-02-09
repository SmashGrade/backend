package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"gorm.io/gorm/clause"
)

type CourseRepository struct {
	Provider db.Provider
}

func NewCourseRepository(provider db.Provider) *CourseRepository {
	return &CourseRepository{
		Provider: provider,
	}
}

func (r *CourseRepository) Create(course models.Course) error {
	result := r.Provider.DB().Create(&course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CourseRepository) Update(course models.Course) error {
	result := r.Provider.DB().Model(&course).Updates(&course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CourseRepository) Find(course models.Course) ([]models.Course, error) {
	var courses []models.Course
	result := r.Provider.DB().Preload(clause.Associations).Where(&course).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func (r *CourseRepository) GetAll() ([]models.Course, error) {
	var courses []models.Course
	result := r.Provider.DB().Preload(clause.Associations).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func (r *CourseRepository) Get(id, version uint) (models.Course, error) {
	var course models.Course
	result := r.Provider.DB().Preload(clause.Associations).Where("id = ? AND version = ?", id, version).First(&course)
	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}

func (r *CourseRepository) Delete(id, version uint) error {
	result := r.Provider.DB().Where("id = ? AND version = ?", id, version).Delete(&models.Course{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
