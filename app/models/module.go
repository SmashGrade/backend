package models

type Module struct {
	VersionedBasemodel
	State            State          `json:"state"`
	StudyStageId     uint           `json:"-"`
	StudyStage       StudyStage     `gorm:"foreignKey:StudyStageId" json:"studyStage"`
	EvaluationTypeID uint           `json:"-"`
	EvaluationType   Evaluationtype `json:"evaluationType"`
	Description      string         `json:"description"`
	Number           string         `json:"number"`
	Curriculums      []*Curriculum  `gorm:"many2many:curriculum_module_assignment;" json:"curriculums"`
	Courses          []*Course      `gorm:"many2many:module_course_assignment;" json:"courses"`
}
