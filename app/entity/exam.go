package entity

type Exam struct {
	basemodel
	Course      Course
	Examtype    Examtype
	Description string
	Weighting   float64
}
