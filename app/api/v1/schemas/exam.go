package schemas

type ExamRes struct {
	Id          int64   `json:"id"`
	Description string  `json:"description"`
	Weight      float64 `json:"weight"`
	Type        string  `json:"string"`
}

type ExamReq struct {
	Description string  `json:"description"`
	Weight      float64 `json:"weight"`
	Type        string  `json:"type"`
}

type ExamReqStudent struct {
	Grade      float64 `json:"grade"`
	Type       string  `json:"type"`
	StudentRef int64   `json:"studentRef"`
}

type CourseExamTeacher struct {
	Id          int64      `json:"id"`
	Description string     `json:"description"`
	Weight      float64    `json:"weight"`
	Type        string     `json:"type"`
	AvgGrades   []GradType `json:"avgGrades"`
}

type CourseExamStudent struct {
	Id          int64    `json:"id"`
	Description string   `json:"description"`
	Weight      float64  `json:"weight"`
	Type        string   `json:"type"`
	Grade       GradeRes `json:"grade"`
}
