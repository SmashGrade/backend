package models

type Exam struct {
	Basemodel
	CourseID      uint     `json:"courseId"`
	CourseVersion uint     `json:"courseVersion"`
	Course        Course   `gorm:"foreignkey:CourseID,CourseVersion;association_foreignkey:ID,Version"`
	ExamtypeID    uint     `json:"-"`
	Examtype      Examtype `json:"examtype"`
	Description   string   `json:"description"`
	Weighting     uint     `json:"weighting"`
}
