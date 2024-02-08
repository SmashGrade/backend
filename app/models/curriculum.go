package models

import "time"

type Curriculum struct {
	Basemodel
	StartValidity    time.Time      `gorm:"primarykey" json:"startValidity"`
	FocusID          uint           `json:"-"`
	Focus            Focus          `json:"focus"`
	CurriculumtypeID uint           `json:"-"`
	Curriculumtype   Curriculumtype `json:"curriculumType"`
	StateID          uint           `json:"-"`
	State            State          `json:"state"`
	EndValidity      time.Time      `json:"endValidity"`
	Description      string         `json:"description"`
	Modules          []*Module      `gorm:"many2many:curriculum_module_assignment;" json:"modules"`
}
