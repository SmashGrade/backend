package entity

type Course struct {
	basemodel
	Version         uint `gorm:"primarykey"`
	Description     string
	Number          string    // this is a short identifier
	Modules         []*Module `gorm:"many2many:module_course_assignment;"`
	TeachedBy       []*User   `gorm:"many2many:course_teacher;"`
	SelectedCourses []SelectedCourse
}
