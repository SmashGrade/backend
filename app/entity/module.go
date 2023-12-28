package entity

type Module struct {
	basemodel
	Version          uint `gorm:"primarykey"`
	State            State
	EvaluationTypeID uint
	Description      string
	Number           string        // this is the short identifier of a module not a real number
	Curriculums      []*Curriculum `gorm:"many2many:curriculum_module_assignment;"`
	Courses          []*Course     `gorm:"many2many:module_course_assignment;"`
}
