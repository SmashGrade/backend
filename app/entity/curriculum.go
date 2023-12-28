package entity

import "time"

type Curriculum struct {
	Basemodel
	StartValidity    time.Time `gorm:"primarykey"`
	FocusID          uint      // belongs to
	Focus            Focus
	CurriculumtypeID uint
	Curriculumtype   Curriculumtype
	StateID          uint
	State            State
	EndValidity      time.Time
	Description      string
	Modules          []*Module `gorm:"many2many:curriculum_module_assignment;"`
}
