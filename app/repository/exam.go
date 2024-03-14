package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"gorm.io/gorm/clause"
)

type ExamRepository struct {
	*BaseRepository
}

func NewExamRepository(provider db.Provider) *ExamRepository {
	return &ExamRepository{
		BaseRepository: NewBaseRepository(provider, models.Exam{}),
	}
}

/*
Get exams by providing the courseID and courseVersion

Usage only for Exam

	exams, err := repository.GetFromCourse(1, 2)
*/
func (r *ExamRepository) GetFromCourse(courseID uint, courseVersion uint) ([]models.Exam, error) {
	var exams []models.Exam

	result := r.Provider.DB().Preload(clause.Associations).Where("course_id = ? AND course_version = ?", courseID, courseVersion).Find(&exams)
	if result.Error != nil {
		return nil, result.Error
	}
	return exams, nil
}
