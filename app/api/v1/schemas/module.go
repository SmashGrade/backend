package schemas

type Module struct {
	Id          int64  `json:"id"`
	Version     int64  `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
	IsActiv     bool   `json:"isActiv"`
}

type ModuleRes struct {
	Id                int64             `json:"id"`
	Version           int64             `json:"version"`
	Description       string            `json:"description"`
	Number            string            `json:"number"`
	IsActiv           bool              `json:"isActiv"`
	StudyStage        StudyStage        `json:"studyStage"`
	ValuationCategory ValuationCategory `json:"valuationCategory"`
	Courses           []CourseRes       `json:"courses"`
}

type ModuleResStudent struct {
	Id                int64              `json:"id"`
	Version           int64              `json:"version"`
	Description       string             `json:"description"`
	Number            string             `json:"number"`
	IsActiv           bool               `json:"isActiv"`
	GradesPerType     []GradType         `json:"gradesPerType"`
	StudyStage        StudyStage         `json:"studyStage"`
	ValuationCategory ValuationCategory  `json:"valuationCategory"`
	Courses           []CourseResStudent `json:"courses"`
}

type ModuleResTeacher struct {
	Id                int64              `json:"id"`
	Version           int64              `json:"version"`
	Description       string             `json:"description"`
	Number            string             `json:"number"`
	IsActiv           bool               `json:"isActiv"`
	AvgGradesPerType  []GradType         `json:"avgGradesPerType"`
	StudyStage        StudyStage         `json:"studyStage"`
	ValuationCategory ValuationCategory  `json:"valuationCategory"`
	Courses           []CourseResTeacher `json:"courses"`
}

type ModuleReq struct {
	Description       string            `json:"description"`
	Number            string            `json:"number"`
	IsActiv           bool              `json:"isActiv"`
	ValuationCategory ValuationCategory `json:"valuationCategory"`
	CoursesRef        []int64           `json:"coursesRef"`
}

type ModuleFilter struct {
	StudyStages []StudyStage `json:"studyStages"`
}
