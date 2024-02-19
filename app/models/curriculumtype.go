package models

type Curriculumtype struct {
	Basemodel
	Description   string `json:"description"`
	DurationYears uint   `json:"durationYears"`
}
