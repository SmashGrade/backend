package models

import (
	"time"

	"gorm.io/gorm"
)

// Grade conversion from raw input in the exam evaluation in a gradetype
type Conversion struct {
	ID        uint           `gorm:"primarykey;autoincrement:false" json:"id"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted,omitempty"`

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
