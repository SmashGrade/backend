package entity

import "time"

// Entered result grade for a specific examn in a course by a user
// OriginalValue saved in raw string input, see conversion for numeric
type ExamEvaluation struct {
	Basemodel                                   // this contains an ID, maybe problematic
	RegisteredByID               uint           `gorm:"primarykey"`
	RegisteredBy                 User           `gorm:"foreignKey:RegisteredByID"`
	SelectedCourseUserID         uint           `gorm:"primarykey"`
	SelectedCourseCourseID       uint           `gorm:"primarykey"`
	SelectedCourseCourseVersion  uint           `gorm:"primarykey"`
	SelectedCourseClassStartyear time.Time      `gorm:"primarykey"`
	SelectedCourse               SelectedCourse `gorm:"foreignKey:SelectedCourseUserID,SelectedCourseCourseID,SelectedCourseCourseVersion,SelectedCourseClassStartyear"`
	ExamID                       uint           `gorm:"primarykey"`
	Exam                         Exam
	OriginalValue                string
	EntryDate                    time.Time
}
