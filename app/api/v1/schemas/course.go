package schemas

type CourseRes struct {
	Id          int64  `json:"id"`
	Version     int64  `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
}

type CourseResStudent struct {
	Id          int64               `json:"id"`
	Version     int64               `json:"version"`
	Description string              `json:"description"`
	Number      string              `json:"number"`
	Exams       []CourseExamStudent `json:"exams"`
}

type CourseResTeacher struct {
	Id          int64               `json:"id"`
	Version     int64               `json:"version"`
	Description string              `json:"description"`
	Number      string              `json:"number"`
	Exams       []CourseExamTeacher `json:"exams"`
}

type CourseReqPut struct {
	Version     int64  `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
	ModuleRef   int64  `json:"integer"`
}

type CourseReqPost struct {
	Description string `json:"description" validate:"required"`
	Number      string `json:"number"`
	ModuleRef   int64  `json:"moduleRef"`
}

type CourseFilter struct {
	Modules  []Module `json:"modules"`
	Versions []int64  `json:"versions"`
}
