package models

type MetaCourse struct {
	Teachers  []User     `json:"teachers"`
	Modules   []Module   `json:"modules"`
	Examtypes []Examtype `json:"examtypes"`
}

type MetaModules struct {
	Evaluationtypes []Evaluationtype `json:"evaluationtype"`
	Curriculums     []Curriculum     `json:"curriculums"`
	Curriculumstype []Curriculumtype `json:"curriculumstype"`
	Courses         []Course         `json:"courses"`
}
