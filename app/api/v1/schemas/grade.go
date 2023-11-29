package schemas

type GradType struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Grade       float64 `json:"grade"`
}

type GradeRes struct {
	Id            int        `json:"id"`
	Date          string     `json:"date"`
	GradesPerType []GradType `json:"gradesPerType"`
}
