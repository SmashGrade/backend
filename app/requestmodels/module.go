package requestmodels

type RefModule struct {
	RefVersioned
	State          RefId          `json:"state"`
	StudyStage     RefId          `json:"studyStage"`
	EvaluationType RefId          `json:"evaluationType"`
	Description    string         `json:"description"`
	Number         string         `json:"number"`
	Curriculums    []RefTimed     `json:"curriculums"`
	Courses        []RefVersioned `json:"courses"`
}
