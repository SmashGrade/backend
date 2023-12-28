package entity

// Grade conversion from raw input in the exam evaluation in a gradetype
type Conversion struct {
	basemodel
	ExamEvaluation ExamEvaluation `gorm:"primarykey"`
	Gradetype      Gradetype      `gorm:"primarykey"`
	Value          float64
}
