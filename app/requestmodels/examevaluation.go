package requestmodels

import "time"

// Entered result grade for a specific exam in a course by a user
// OriginalValue saved in raw string input, see conversion for numeric
type RefExamEvaluation struct {
	EvaluationID                 uint      `json:"evaluationID"`
	RegisteredByID               uint      `json:"registeredByID"`
	SelectedCourseUserID         uint      `json:"selectedCourseUserID"`
	SelectedCourseCourseID       uint      `json:"selectedCourseCourseID"`
	SelectedCourseCourseVersion  uint      `json:"selectedCourseCourseVersion"`
	SelectedCourseClassStartyear time.Time `json:"selectedCourseClassStartyear"`
	ExamID                       uint      `json:"examID"`
	OriginalValue                string    `json:"originalValue"`
	OrignialGradeTypeID          uint      `json:"orignialGradeTypeID"`
	EntryDate                    time.Time `json:"entryDate"`
}
