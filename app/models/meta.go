package models

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
	Curriculumstype []Curriculumtype `json:"curriculumtypes"`
	Users           []User           `json:"users"`
}
