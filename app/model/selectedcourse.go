package entity

import "time"

type SelectedCourse struct {
	UserID         uint      `gorm:"primarykey"`
	CourseID       uint      `gorm:"primarykey"`
	CourseVersion  uint      `gorm:"primarykey"`
	ClassStartyear time.Time `gorm:"primarykey"`
	Dispensed      bool
}
