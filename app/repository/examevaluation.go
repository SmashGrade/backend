package repository

import (
	"time"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type ExamEvaluationRepository struct {
	*BaseRepository
}

func NewExamEvaluationRepository(provider db.Provider) *ExamEvaluationRepository {
	return &ExamEvaluationRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}

func (r *ExamEvaluationRepository) GetExamEvaluation(
	id,
	registeredById,
	selectedCourseUserId,
	selectedCourseCourseId,
	selectedCourseCourseVersion,
	examId uint,
	selectedCourseClassStartyear time.Time) (examEvaluation models.ExamEvaluation, err error) {
	result := r.Provider.DB().Where("id = ? AND registered_by_id = ? AND selected_course_user_id = ? AND selected_course_course_id = ? AND selected_course_course_version = ? AND selected_course_class_startyear = ? exam_id = ?",
		id, registeredById, selectedCourseUserId, selectedCourseCourseId, selectedCourseCourseVersion, selectedCourseClassStartyear, examId).First(&examEvaluation)
	err = result.Error
	return
}

func (r *ExamEvaluationRepository) DeleteExamEvaluation(id,
	registeredById,
	selectedCourseUserId,
	selectedCourseCourseId,
	selectedCourseCourseVersion,
	examId uint,
	selectedCourseClassStartyear time.Time) error {
	return r.Provider.DB().Where("id = ? AND registered_by_id = ? AND selected_course_user_id = ? AND selected_course_course_id = ? AND selected_course_course_version = ? AND selected_course_class_startyear = ? exam_id = ?",
		id, registeredById, selectedCourseUserId, selectedCourseCourseId, selectedCourseCourseVersion, selectedCourseClassStartyear, examId).Delete(&models.ExamEvaluation{}).Error

}
