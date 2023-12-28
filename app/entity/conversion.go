package entity

import "time"

// Grade conversion from raw input in the exam evaluation in a gradetype
type Conversion struct {
	Basemodel

	EERegisteredByID               uint      `gorm:"primarykey"`
	EESelectedCourseUserID         uint      `gorm:"primarykey"`
	EESelectedCourseCourseID       uint      `gorm:"primarykey"`
	EESelectedCourseCourseVersion  uint      `gorm:"primarykey"`
	EESelectedCourseClassStartyear time.Time `gorm:"primarykey"`
	EEExamID                       uint      `gorm:"primarykey"`

	ExamEvaluation ExamEvaluation `gorm:"foreignKey:EERegisteredByID,EESelectedCourseUserID,EESelectedCourseCourseID,EESelectedCourseCourseVersion,EESelectedCourseClassStartyear,EEExamID;References:RegisteredByID,SelectedCourseUserID,SelectedCourseCourseID,SelectedCourseCourseVersion,SelectedCourseClassStartyear,ExamID"`
	GradetypeID    uint           `gorm:"primarykey"`
	Gradetype      Gradetype
	Value          float64
}
