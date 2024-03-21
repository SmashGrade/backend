package requestmodels

type RefCurriculum struct {
	RefTimed
	FocusID          RefId          `json:"focus"`
	CurriculumtypeID RefId          `json:"curriculumType"`
	StateID          RefId          `json:"state"`
	EndValidity      string         `json:"endValidity" example:"02.01.2006"`
	Description      string         `json:"description" example:"Softwareentwicklung"`
	Modules          []RefVersioned `json:"modules"`
}

type RefOnboarding struct {
	StartDate    string `json:"startDate"`
	CurriculumID uint   `json:"curriculumId"`
}
