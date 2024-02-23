package models

import "time"

// Entered result grade for a specific exam in a course by a user
// OriginalValue saved in raw string input, see conversion for numeric
type ExamEvaluation struct {
	Basemodel
	RegisteredByID               uint           `gorm:"primarykey" json:"registeredByID"`
	RegisteredBy                 User           `gorm:"foreignKey:RegisteredByID" json:"registeredBy"`
	SelectedCourseUserID         uint           `gorm:"primarykey" json:"selectedCourseUserID"`
	SelectedCourseCourseID       uint           `gorm:"primarykey" json:"selectedCourseCourseID"`
	SelectedCourseCourseVersion  uint           `gorm:"primarykey" json:"selectedCourseCourseVersion"`
	SelectedCourseClassStartyear time.Time      `gorm:"primarykey" json:"selectedCourseClassStartyear"`
	SelectedCourse               SelectedCourse `gorm:"foreignKey:SelectedCourseUserID,SelectedCourseCourseID,SelectedCourseCourseVersion,SelectedCourseClassStartyear" json:"selectedCourse"`
	ExamID                       uint           `gorm:"primarykey" json:"examID"`
	Exam                         Exam           `json:"exam"`
	OriginalValue                string         `json:"originalValue"`
	EntryDate                    time.Time      `json:"entryDate"`
}
