package entity

import "time"

type User struct {
	Basemodel
	ClassStartyear  time.Time
	Name            string
	Email           string
	Fields          []*Field  `gorm:"many2many:fieldmanager;"`
	Roles           []*Role   `gorm:"many2many:user_has_role;"`
	TeachesCourses  []*Course `gorm:"many2many:course_teacher;"`
	SelectedCourses []SelectedCourse
}
