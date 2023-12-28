package entity

type Exam struct {
	Basemodel
	CourseID    uint
	Course      Course
	ExamtypeID  uint
	Examtype    Examtype
	Description string
	Weighting   float64
}
