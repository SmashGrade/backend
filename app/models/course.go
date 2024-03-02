package models

type Course struct {
	VersionedBasemodel
	Description     string           `json:"description"`
	Number          string           `json:"number"`
	Modules         []*Module        `gorm:"many2many:module_course_assignment;" json:"modules"`
	TeachedBy       []*User          `gorm:"many2many:course_teacher;" json:"teachedBy"`
	SelectedCourses []SelectedCourse `json:"selectedCourses"`
	Exams           []*Exam          `json:"exams"`
}
