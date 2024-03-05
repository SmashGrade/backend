package requestmodels

type RefCurriculum struct {
	RefTimed
	FocusID          RefId          `json:"focus"`
	CurriculumtypeID RefId          `json:"curriculumType"`
	StateID          RefId          `json:"state"`
	EndValidity      string         `json:"endValidity"`
	Description      string         `json:"description"`
	Modules          []RefVersioned `json:"modules"`
}
