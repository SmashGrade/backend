package models

import "time"

type SelectedCourse struct {
	UserID         uint      `gorm:"primarykey" json:"userId"`
	CourseID       uint      `gorm:"primarykey" json:"courseId"`
	CourseVersion  uint      `gorm:"primarykey" json:"courseVersion"`
	ClassStartyear time.Time `gorm:"primarykey" json:"classStartYear"`
	Dispensed      bool      `json:"dispensed"`
}
