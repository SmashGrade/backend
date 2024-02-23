package entity

type Exam struct {
	Basemodel
	CourseID      uint
	CourseVersion uint
	Course        Course `gorm:"foreignkey:CourseID,CourseVersion;association_foreignkey:ID,Version"`
	ExamtypeID    uint
	Examtype      Examtype
	Description   string
	Weighting     float64
}
