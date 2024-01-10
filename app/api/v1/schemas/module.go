package schemas

type Module struct {
	Id          uint   `json:"id"`
	Version     uint   `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
	IsActiv     bool   `json:"isActive"`
}

type ModuleRes struct {
	Id                uint              `json:"id"`
	Version           uint              `json:"version"`
	Description       string            `json:"description"`
	Number            string            `json:"number"`
	IsActiv           bool              `json:"isActive"`
	StudyStage        StudyStage        `json:"studyStage"`
	ValuationCategory ValuationCategory `json:"valuationCategory"`
	Courses           []CoursesRes      `json:"courses"`
}

type ModuleResStudent struct {
	Id                uint               `json:"id"`
	Version           uint               `json:"version"`
	Description       string             `json:"description"`
	Number            string             `json:"number"`
	IsActiv           bool               `json:"isActive"`
	GradesPerType     []GradType         `json:"gradesPerType"`
	StudyStage        StudyStage         `json:"studyStage"`
	ValuationCategory ValuationCategory  `json:"valuationCategory"`
	Courses           []CourseResStudent `json:"courses"`
}

type ModuleResTeacher struct {
	Id                uint               `json:"id"`
	Version           uint               `json:"version"`
	Description       string             `json:"description"`
	Number            string             `json:"number"`
	IsActiv           bool               `json:"isActive"`
	AvgGradesPerType  []GradType         `json:"avgGradesPerType"`
	StudyStage        StudyStage         `json:"studyStage"`
	ValuationCategory ValuationCategory  `json:"valuationCategory"`
	Courses           []CourseResTeacher `json:"courses"`
}

type ModuleReq struct {
	Description       string            `json:"description"`
	Number            string            `json:"number"`
	IsActiv           bool              `json:"isActive"`
	ValuationCategory ValuationCategory `json:"valuationCategory"`
	CoursesRef        []uint            `json:"coursesRef"`
}

type ModuleFilter struct {
	StudyStages []StudyStage `json:"studyStages"`
}

type ModuleRef struct {
	Id      uint `json:"id"`
	Version uint `json:"version"`
}
