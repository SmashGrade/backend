package entity

import "time"

type Module struct {
	ID               uint `gorm:"primarykey;autoIncrement:false"` // cant autoincrement else update with version wont work
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Version          uint `gorm:"primarykey"`
	State            State
	StudyStageId     uint
	StudyStage       StudyStage `gorm:"foreignKey:StudyStageId"`
	EvaluationTypeID uint
	Description      string
	Number           string        // this is the short identifier of a module not a real number
	Curriculums      []*Curriculum `gorm:"many2many:curriculum_module_assignment;"`
	Courses          []*Course     `gorm:"many2many:module_course_assignment;"`
}