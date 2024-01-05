package schemas

type CoursesRes struct {
	Id          int    `json:"id"`
	Version     int    `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
	Versions    []int  `json:"versions"`
}

type CourseRes struct {
	Id          int       `json:"id"`
	Version     int       `json:"version"`
	Description string    `json:"description"`
	Number      string    `json:"number"`
	Versions    []int     `json:"versions"`
	Modules     []Module  `json:"modules"`
	Exams       []ExamRes `json:"exams"`
	Teachers    []Teacher `json:"teacher"`
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
	Version     int       `json:"version"`
	Description string    `json:"description"`
	Number      string    `json:"number"`
	ModuleRef   []int     `json:"moduleRef"`
	TeacherRef  []int     `json:"teacherRef"`
	Exams       []ExamRes `json:"exams"`
}

type CourseReqPost struct {
	Description string      `json:"description" validate:"required"`
	Number      string      `json:"number"`
	ModuleRef   []ModuleRef `json:"moduleRef"`
	TeacherRef  []int       `json:"teacherRef"`
	Exams       []ExamRes   `json:"exams"`
}

type CourseFilter struct {
	Modules  []Module  `json:"modules"`
	Teachers []Teacher `json:"teachers"`
}
