package schemas

type CoursesRes struct {
	Id          uint   `json:"id"`
	Version     uint   `json:"version"`
	Description string `json:"description"`
	Number      string `json:"number"`
	Versions    []uint `json:"versions"`
}

type CourseRes struct {
	Id          uint      `json:"id"`
	Version     uint      `json:"version"`
	Description string    `json:"description"`
	Number      string    `json:"number"`
	Versions    []uint    `json:"versions"`
	Modules     []Module  `json:"modules"`
	Exams       []ExamRes `json:"exams"`
	Teachers    []Teacher `json:"teacher"`
}

type CourseResStudent struct {
	Id          uint                `json:"id"`
	Version     uint                `json:"version"`
	Description string              `json:"description"`
	Number      string              `json:"number"`
	Exams       []CourseExamStudent `json:"exams"`
}

type CourseResTeacher struct {
	Id          uint                `json:"id"`
	Version     uint                `json:"version"`
	Description string              `json:"description"`
	Number      string              `json:"number"`
	Exams       []CourseExamTeacher `json:"exams"`
}

type CourseReqPut struct {
	Version     uint      `json:"version"`
	Description string    `json:"description"`
	Number      string    `json:"number"`
	ModuleRef   []uint    `json:"moduleRef"`
	TeacherRef  []uint    `json:"teacherRef"`
	Exams       []ExamRes `json:"exams"`
}

type CourseReqPost struct {
	Description string      `json:"description" validate:"required"`
	Number      string      `json:"number"`
	ModuleRef   []ModuleRef `json:"moduleRef"`
	TeacherRef  []uint      `json:"teacherRef"`
	Exams       []ExamRes   `json:"exams"`
}

type CourseFilter struct {
	Modules  []Module  `json:"modules"`
	Teachers []Teacher `json:"teachers"`
}
