package requestmodels

type RefExam struct {
	RefId
	Course      RefVersioned `json:"course"`
	ExamType    RefId        `json:"examtype"`
	Description string       `json:"description"`
	Weighting   uint         `json:"weighting"`
}
