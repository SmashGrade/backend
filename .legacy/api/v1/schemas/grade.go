package schemas

type GradType struct {
	Id          uint    `json:"id"`
	Description string  `json:"description"`
	Grade       float64 `json:"grade"`
}

type GradeRes struct {
	Id            uint       `json:"id"`
	Date          string     `json:"date"`
	GradesPerType []GradType `json:"gradesPerType"`
}
