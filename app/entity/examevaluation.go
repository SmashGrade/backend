package entity

import "time"

// Entered result grade for a specific examn in a course by a user
// OriginalValue saved in raw string input, see conversion for numeric
type ExamEvaluation struct {
	basemodel                     // this contains an ID, maybe problematic
	RegisteredBy   User           `gorm:"primarykey"`
	SelectedCourse SelectedCourse `gorm:"primarykey"`
	Exam           Exam           `gorm:"primarykey"`
	OriginalValue  string
	EntryDate      time.Time
}
