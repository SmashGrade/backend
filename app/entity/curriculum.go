package entity

import "time"

type Curriculum struct {
	basemodel
	StartValidity  time.Time `gorm:"primarykey"`
	FocusID        uint      // belongs to
	Focus          Focus
	Curriculumtype Curriculumtype
	State          State
	EndValidity    time.Time
	Description    string
	Modules        []*Module `gorm:"many2many:curriculum_module_assignment;"`
}