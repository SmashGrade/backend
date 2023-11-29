package schemas

type CourseRes struct {
	Id          int    `json:"id"`
	Version     int    `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
}

type CourseResStudent struct {
	Id          int                 `json:"id"`
	Version     int                 `json:"version"`
	Description string              `json:"description"`
	Number      string              `json:"number"`
	Exams       []CourseExamStudent `json:"exams"`
}

type CourseResTeacher struct {
	Id          int                 `json:"id"`
	Version     int                 `json:"version"`
	Description string              `json:"description"`
	Number      string              `json:"number"`
	Exams       []CourseExamTeacher `json:"exams"`
}

type CourseReqPut struct {
	Version     int    `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
	ModuleRef   int    `json:"integer"`
}

type CourseReqPost struct {
	Description string `json:"description" validate:"required"`
	Number      string `json:"number"`
	ModuleRef   int    `json:"moduleRef"`
}

type CourseFilter struct {
	Modules  []Module `json:"modules"`
	Versions []int    `json:"versions"`
}
