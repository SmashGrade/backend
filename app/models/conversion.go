package models

import "time"

// Grade conversion from raw input in the exam evaluation in a gradetype
type Conversion struct {
	Basemodel

	EERegisteredByID               uint      `gorm:"primarykey" json:"EERegisteredByID"`
	EESelectedCourseUserID         uint      `gorm:"primarykey" json:"EESelectedCourseUserID"`
	EESelectedCourseCourseID       uint      `gorm:"primarykey" json:"EESelectedCourseCourseID"`
	EESelectedCourseCourseVersion  uint      `gorm:"primarykey" json:"EESelectedCourseCourseVersion"`
	EESelectedCourseClassStartyear time.Time `gorm:"primarykey" json:"EESelectedCourseClassStartyear"`
	EEExamID                       uint      `gorm:"primarykey" json:"EEExamID"`

	ExamEvaluation ExamEvaluation `gorm:"foreignKey:EERegisteredByID,EESelectedCourseUserID,EESelectedCourseCourseID,EESelectedCourseCourseVersion,EESelectedCourseClassStartyear,EEExamID;References:RegisteredByID,SelectedCourseUserID,SelectedCourseCourseID,SelectedCourseCourseVersion,SelectedCourseClassStartyear,ExamID" json:"ExamEvaluation"`
	GradetypeID    uint           `gorm:"primarykey" json:"GradetypeID"`
	Gradetype      Gradetype      `json:"Gradetype"`
	Value          float64        `json:"Value"`
}
