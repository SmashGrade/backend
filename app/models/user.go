package models

import (
	"time"
)

type User struct {
	Basemodel
	ClassStartyear  time.Time        `json:"classStartYear"`
	Name            string           `json:"name"`
	Email           string           `gorm:"unique" json:"email"`
	Fields          []*Field         `gorm:"many2many:fieldmanager;" json:"fields"`
	Roles           []*Role          `gorm:"many2many:user_has_role;" json:"roles"`
	TeachesCourses  []*Course        `gorm:"many2many:course_teacher;" json:"teachesCourses"`
	SelectedCourses []SelectedCourse `json:"selectedCourses"`
	CurriculumID    uint             `json:"curriculumId"`
}

// Returns true if the user has the role with the given id
func (u User) HasRole(roleId uint) bool {
	for i := range u.Roles {
		if u.Roles[i].ID == roleId {
			return true
		}
	}
	return false
}

// Returns true if the user has any role
func (u User) HasAnyRole() bool {
	return len(u.Roles) > 0
}
