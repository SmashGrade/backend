package entity

import "time"

type Curriculum struct {
	Basemodel
	StartValidity    time.Time `gorm:"primarykey"`
	FocusID          uint      // belongs to
	Focus            Focus     `json:"-"`
	CurriculumtypeID uint
	Curriculumtype   Curriculumtype `json:"-"`
	StateID          uint
	State            State
	EndValidity      time.Time
	Description      string
	Modules          []*Module `gorm:"many2many:curriculum_module_assignment;"`
}
