package schemas

type GradType struct {
	Id          int64   `json:"id"`
	Description string  `json:"description"`
	Grade       float64 `json:"grade"`
}

type GradeRes struct {
	Id            int64      `json:"id"`
	Date          string     `json:"date"`
	GradesPerType []GradType `json:"gradesPerType"`
}
