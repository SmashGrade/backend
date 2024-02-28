package requestmodels

type RefCourse struct {
	RefVersioned
	Description     string              `json:"description"`
	Number          string              `json:"number"`
	Modules         []RefVersioned      `json:"modules"`
	TeachedBy       []RefId             `json:"teachedBy"`
	SelectedCourses []RefSelectedCourse `json:"selectedCourses"`
	Exams           []RefId             `json:"exams"`
}
