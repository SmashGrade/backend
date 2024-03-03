package models

import "time"

type MetaCourse struct {
	Teachers  []User     `json:"teachers"`
	Modules   []Module   `json:"modules"`
	Examtypes []Examtype `json:"examtypes"`
}

type MetaModules struct {
	Evaluationtypes []Evaluationtype `json:"evaluationtype"`
	Curriculums     []Curriculum     `json:"curriculums"`
	Curriculumtypes []Curriculumtype `json:"curriculumtypes"`
	Courses         []Course         `json:"courses"`
}

// to fill out selection fields in create/edit view
// contains: all focus (Fachrichtung), all fields (Schwerpunkt), all curriculumtypes, all users
type MetaCurriculums struct {
	Focuses         []Focus          `json:"focuses"`
	Fields          []Field          `json:"fields"`
	Curriculumtypes []Curriculumtype `json:"curriculumtypes"`
	Teachers        []User           `json:"teachers"`
}

// list of courses teached by current user with modules and study stage, list of all users
type TeacherCourses struct {
}

// chosen curriculum with start year and curriculum type
type StudentCurriculums struct {
	StartYear time.Time // taken from the user or selectedcourse
	// Curriculumtype Curriculumtype // this is already in the curriculum
	Curriculum Curriculum
}
