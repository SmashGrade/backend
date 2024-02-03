package entity

import "time"

type Course struct {
	ID              uint `gorm:"primarykey;autoIncrement:false"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Version         uint `gorm:"primarykey"` //;autoIncrement:false"`
	Description     string
	Number          string    // this is a short identifier
	Modules         []*Module `gorm:"many2many:module_course_assignment;"`
	TeachedBy       []*User   `gorm:"many2many:course_teacher;"`
	SelectedCourses []SelectedCourse
}
