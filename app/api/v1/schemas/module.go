package schemas

type Module struct {
	Id          int64  `json:"id"`
	Version     int64  `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
	IsActiv     bool   `json:"isActive"`
}

type ModuleRes struct {
	Id                int64             `json:"id"`
	Version           int64             `json:"version"`
	Description       string            `json:"description"`
	Number            string            `json:"number"`
	IsActiv           bool              `json:"isActive"`
	StudyStage        StudyStage        `json:"studyStage"`
	ValuationCategory ValuationCategory `json:"valuationCategory"`
	Courses           []CourseRes       `json:"courses"`
}

type ModuleResStudent struct {
	Id                int                `json:"id"`
	Version           int                `json:"version"`
	Description       string             `json:"description"`
	Number            string             `json:"number"`
	IsActiv           bool               `json:"isActive"`
	GradesPerType     []GradType         `json:"gradesPerType"`
	StudyStage        StudyStage         `json:"studyStage"`
	ValuationCategory ValuationCategory  `json:"valuationCategory"`
	Courses           []CourseResStudent `json:"courses"`
}

type ModuleResTeacher struct {
	Id                int                `json:"id"`
	Version           int                `json:"version"`
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
	CoursesRef        []int             `json:"coursesRef"`
}

type ModuleFilter struct {
	StudyStages []StudyStage `json:"studyStages"`
}

type ModuleRef struct {
	Id      uint `json:"id"`
	Version uint `json:"version"`
}
